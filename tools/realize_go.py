import os
import re
import json

# 读取配置文件
with open('config.json', 'r', encoding='utf-8') as config_file:
    config = json.load(config_file)

# 从配置文件中获取配置信息
PROTO_DIR = config['proto_dir']
PROTO_DIRS = [os.path.join(PROTO_DIR, d) for d in config['proto_dirs']]
GOFILE_DIRECTORIES = [os.path.join(PROTO_DIR, d) for d in config['gofile_directory']]
FILTERED_FILES = config['filtered_files']
NODE_TO_GO_FILE = config['node_to_go_file']
MESSAGE_ID_OUTPUT = config['message_id_output']

def extract_messages(proto_dirs):
    """
    提取所有未被注释的 message 及其注释。
    返回一个字典，键为节点名，值为消息列表，每个消息为元组 (message_name, comment)
    """
    messages = {}
    message_pattern = re.compile(r'^\s*message\s+(\w+)\s*\{')

    for proto_dir in proto_dirs:
        node_name = os.path.basename(proto_dir)
        if node_name not in messages:
            messages[node_name] = []

        for root, _, files in os.walk(proto_dir):
            for file in files:
                if file.endswith('.proto'):
                    # 检查文件是否在过滤列表中
                    if file in FILTERED_FILES:
                        continue  # 跳过被过滤的文件

                    path = os.path.join(root, file)
                    try:
                        with open(path, 'r', encoding='utf-8') as f:
                            lines = f.readlines()
                            for i, line in enumerate(lines):
                                msg_match = message_pattern.match(line)
                                if msg_match:
                                    msg_name = msg_match.group(1)
                                    
                                    # 检查当前行是否被单行注释 '//' 注释掉
                                    if re.match(r'^\s*//', line):
                                        continue  # 跳过被注释的 message

                                    # 收集上方的连续单行注释
                                    comments = []
                                    j = i - 1
                                    while j >= 0:
                                        comment_match = re.match(r'^\s*//\s?(.*)', lines[j])
                                        if comment_match:
                                            comments.insert(0, comment_match.group(1).strip())
                                            j -= 1
                                        else:
                                            break
                                    full_comment = '\n'.join(comments) if comments else ''
                                    messages[node_name].append((msg_name, full_comment))
                    except Exception as e:
                        print(f'无法读取文件 {path}: {e}')
    return messages

def generate_registration_lines(messages):
    """
    生成注册代码行，不包含注释。
    """
    lines = []
    for msg, _ in messages:
        if msg.endswith('Req') or msg.endswith('Request'):
            method = msg[:-3]  # 去除 'Req' 或 'Request' 后缀
            lines.append(f'\tp.Local().Register(p.{method})')
    return lines

def generate_handler_functions(messages):
    """
    生成处理函数代码，包含注释。
    """
    functions = []
    for msg, comment in messages:
        if msg.endswith('Req') or msg.endswith('Request'):
            method = msg[:-3]  # 去除 'Req' 或 'Request' 后缀
            if comment:
                comment_lines = '// ' + comment.replace('\n', '\n// ')
                functions.append(f'''
    {comment_lines}
    func (p *ActorPlayer) {method}(session *cproto.Session, m *protoMsg.{msg}) {{
    \t// TODO: 实现 {method} 处理逻辑
    }}
    ''')
            else:
                functions.append(f'''
    func (p *ActorPlayer) {method}(session *cproto.Session, m *protoMsg.{msg}) {{
    \t// TODO: 实现 {method} 处理逻辑
    }}
    ''')
    return functions

def update_go_file(go_file, registration_lines, handler_functions):
    try:
        if not os.path.isfile(go_file):
            print(f'文件不存在: {go_file}')
            return

        with open(go_file, 'r', encoding='utf-8') as f:
            content = f.read()

        updated_content = content  # 初始化

        # 匹配 ActorPlayer 的 registerLocalMsg 方法
        register_func_pattern = re.compile(
            r'(func\s+\(p\s+\*ActorPlayer\)\s+registerLocalMsg\s*\(\)\s*\{)([^}]*)\}',
            re.MULTILINE | re.DOTALL
        )
        match = register_func_pattern.search(content)
        if match:
            func_start = match.group(1)
            existing_body = match.group(2).strip()

            # 将现有注册行分割为集合，避免重复注册
            existing_lines = set()
            for line in existing_body.split('\n'):
                line = line.strip()
                if line.startswith('p.Local().Register'):
                    existing_lines.add(line)

            new_lines = []
            for reg_line in registration_lines:
                if reg_line.strip() not in existing_lines:
                    new_lines.append(reg_line)

            if new_lines:
                # 添加新的注册行到现有方法体
                updated_body = existing_body + '\n' + '\n'.join(new_lines)
                updated_func = f'{func_start}\n{updated_body}\n}}'
                updated_content = content.replace(match.group(0), updated_func)
                print('已在现有的 registerLocalMsg 方法中添加注册代码。')
            else:
                print('所有注册消息已存在，未执行更新。')
        else:
            # 如果没有找到 registerLocalMsg 方法，则添加该方法
            func_definition = '\nfunc (p *ActorPlayer) registerLocalMsg() {\n' + '\n'.join(registration_lines) + '\n}\n'
            updated_content += func_definition
            print('已添加 registerLocalMsg 方法并注册消息。')

        # 处理处理函数，避免重复
        handler_pattern = re.compile(r'func\s+\(p\s+\*ActorPlayer\)\s+(\w+)\s*\(', re.MULTILINE)
        existing_handlers = set(re.findall(handler_pattern, content))

        new_functions = []
        for func in handler_functions:
            func_name_match = re.search(r'func\s+\(p\s+\*ActorPlayer\)\s+(\w+)\s*\(', func)
            if func_name_match:
                func_name = func_name_match.group(1)
                if func_name not in existing_handlers:
                    new_functions.append(func)

        if new_functions:
            updated_content += '\n' + '\n'.join(new_functions)
            print('已添加新的处理函数。')
        else:
            print('所有处理函数已存在，未执行添加。')

        # 写回文件
        with open(go_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        print(f'已更新 {go_file} 文件。')

    except OSError as e:
        print(f'文件操作错误: {e}')
    except Exception as e:
        print(f'发生错误: {e}')

def comment_lines_in_file(file_path, target_strings):
    """
    注释文件中包含目标字符串的行。
    """
    with open(file_path, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    modified = False
    new_lines = []
    for line in lines:
        stripped_line = line.strip()
        if any(target in stripped_line for target in target_strings):
            if not stripped_line.startswith("//"):
                new_lines.append("// " + line)
                modified = True
            else:
                new_lines.append(line)
        else:
            new_lines.append(line)

    if modified:
        with open(file_path, 'w', encoding='utf-8') as file:
            file.writelines(new_lines)
        print(f"已修改文件: {file_path}")
    else:
        print(f"未找到需要修改的行: {file_path}")

def traverse_and_modify(directory, target_strings, file_extension=".go"):
    """
    遍历目录下所有指定扩展名的文件，并注释目标字符串所在的行。
    """
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith(file_extension):
                file_path = os.path.join(root, file)
                comment_lines_in_file(file_path, target_strings)

def comment_proto_files():
    """
    注释gofile目录下所有.go文件中包含指定proto语句的行。
    """
    # 设置要注释掉的目标字符串
    targets = [
        'proto "github.com/golang/protobuf/proto"',
        'const _ = proto.ProtoPackageIsVersion4'
    ]

    for gofile_dir in GOFILE_DIRECTORIES:
        traverse_and_modify(gofile_dir, targets)

def update_message_id_file(messages_by_node):
    """
    更新 message_id.json 文件，记录所有以 Req、Resp、Request 和 Response 结尾的消息体及其所属节点。
    """
    message_id_data = {}
    for node_index, (node, messages) in enumerate(messages_by_node.items()):

        node_id =  node_index * 10000 
        count = 0
        for _, (msg, _) in enumerate(messages):
            if msg.endswith('Req') or msg.endswith('Resp') or msg.endswith('Request') or msg.endswith('Response'):
                count += 1
                print(f"{node_id + count}: {msg} ")
                message_id_data[node_id + count] = {
                    "name": msg,
                    "node": node
                }

    # 按照 ID 从小到大排序
    sorted_message_id_data = dict(sorted(message_id_data.items()))

    with open(MESSAGE_ID_OUTPUT, 'w', encoding='utf-8') as f:
        json.dump(sorted_message_id_data, f, indent=4, ensure_ascii=False)
    print(f'已更新 {MESSAGE_ID_OUTPUT} 文件。')

def replace_import_path_in_go_files(gofile_directories):
    """
    替换生成的 .go 文件中的 import 路径
    """
    for gofile_dir in gofile_directories:
        for root, _, files in os.walk(gofile_dir):
            for file in files:
                if file.endswith('.go'):
                    file_path = os.path.join(root, file)
                    with open(file_path, 'r', encoding='utf-8') as f:
                        content = f.read()
                    content = content.replace('pb "/pb"', 'pb "superman/internal/protocol/go_file/common"')
                    with open(file_path, 'w', encoding='utf-8') as f:
                        f.write(content)
                    print(f'已更新 import 路径: {file_path}')

def main():
    # 提取并处理消息注册和处理函数
    messages_by_node = extract_messages(PROTO_DIRS)
    for node, messages in messages_by_node.items():
        unique_messages = list(set(messages))  # 去重
        if not unique_messages:
            print(f'未找到有效 message 在节点: {node}')
            continue

        # 生成注册代码
        registration_lines = generate_registration_lines(unique_messages)

        # 生成处理函数代码
        handler_functions = generate_handler_functions(unique_messages)

        # 获取对应的 Go 文件路径
        go_file = NODE_TO_GO_FILE.get(node)
        if go_file:
            # 更新 Go 文件
            update_go_file(go_file, registration_lines, handler_functions)
    
    # 注释gofile目录下的指定proto语句
    comment_proto_files()

    # 更新 message_id.json 文件
    update_message_id_file(messages_by_node)

    # 替换生成的 .go 文件中的 import 路径
    replace_import_path_in_go_files(GOFILE_DIRECTORIES)

if __name__ == "__main__":
    main()