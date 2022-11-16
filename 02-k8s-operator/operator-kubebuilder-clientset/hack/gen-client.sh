#!/bin/bash
set -e
set -x

# 生成 clientset 代码

# 获取 go module name
go_module=$(go list -m) 
# crd group
group=${GROUP:-"job"} 
# api 版本
api_version=${API_VERSION:-"v1"}

project_dir=$(cd $(dirname ${BASH_SOURCE[0]})/..; pwd) # 项目根目录

# check generate-groups.sh is exist
if [ ! -f "$project_dir/hack/generate-groups.sh" ]; then
    echo "hack/generate-groups.sh is not exist, download"

    wget -O "$project_dir/hack/generate-groups.sh" https://raw.githubusercontent.com/kubernetes/code-generator/master/generate-groups.sh
    chmod +x $project_dir/hack/generate-groups.sh
fi

# 生成 clientset
CLIENTSET_NAME_VERSIONED="$api_version" \
$project_dir/hack/generate-groups.sh client \
    $go_module/pkg $go_module/api "$group:$api_version" --output-base $project_dir/

# 生成的 clientset 的文件夹路径会包含 $go_module/pkg 所以我们需要把这个文件夹复制出来
if [ ! -d "$project_dir/pkg" ];then
    mkdir $project_dir/pkg
fi

rm -rf $project_dir/pkg/clientset
mv -f $project_dir/$go_module/pkg/* $project_dir/pkg/

# 删除不需要的目录
rm -rf $project_dir/$(echo $go_module | cut -d '/' -f 1)
