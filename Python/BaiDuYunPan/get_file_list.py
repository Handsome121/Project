import os
from baidupcs_py.baidupcs import BaiduPCS
from baidupcs_py.commands.sifter import Sifter

# 设置你的百度云账号的用户名和密码
username = "951791665@qq.com"
password = "19950308.zxc"

# 文件下载到本地的路径
local_path = "./download"


def list_all_folders(pcs, folder_path='/'):
    folders = []
    entries = pcs.list(folder_path)
    for entry in entries:
        if entry.is_dir:
            folders.append(entry.path)
            folders.extend(list_all_folders(pcs, entry.path))
    return folders


def download_folder(pcs, remote_folder_path, local_path):
    entries = pcs.list(remote_folder_path)
    for entry in entries:
        remote_path = entry.path
        local_file_path = os.path.join(local_path, os.path.relpath(remote_path, remote_folder_path))

        if entry.is_file:
            print(f"开始下载：{remote_path}")
            os.makedirs(os.path.dirname(local_file_path), exist_ok=True)
            with open(local_file_path, "wb") as f:
                for data in pcs.download(remote_path, sifter=Sifter()):
                    f.write(data)
            print(f"下载完成：{remote_path}")

        elif entry.is_dir:
            download_folder(pcs, remote_path, local_path)


def main():
    pcs = BaiduPCS(username, password)
    # 列出所有文件夹
    all_folders = list_all_folders(pcs)
    print(all_folders)
    # # 打印所有文件夹，供用户选择
    # for index, folder in enumerate(all_folders):
    #     print(f"{index + 1}. {os.path.basename(folder)}")
    #
    # # 让用户选择一个文件夹
    # while True:
    #     folder_index = input("请选择要下载的文件夹序号：")
    #     try:
    #         folder_index = int(folder_index)
    #         if folder_index < 1 or folder_index > len(all_folders):
    #             raise ValueError()
    #         break
    #     except ValueError:
    #         print("无效的输入，请重新输入。")
    #
    # remote_folder_path = all_folders[folder_index - 1]
    # download_folder(pcs, remote_folder_path, local_path)


if __name__ == '__main__':
    main()
