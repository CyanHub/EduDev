document.addEventListener('DOMContentLoaded', function () {
    const downloadForm = document.getElementById('downloadForm');
    const token = localStorage.getItem('token');

    if (!token) {
        Swal.fire({
            icon: 'error',
            title: '错误',
            text: '请先登录'
        });
        return;
    }

    // 暂时注释掉加载文件列表的功能
    /*
    // 加载文件列表
    fetch('http://localhost:4407/files', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('请求失败');
        }
        return response.json();
    })
    .then(data => {
        const filesList = document.getElementById('filesList');
        filesList.innerHTML = '';
        data.forEach(file => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${file.ID}</td>
                <td>${file.Name}</td>
                <td>${file.Owner ? file.Owner.Username : '未知'}</td>
                <td>${file.CreatedAt}</td>
                <td>${file.IsPublic ? '是' : '否'}</td>
            `;
            filesList.appendChild(row);
        });
    })
    .catch(error => {
        Swal.fire({
            icon: 'error',
            title: '错误',
            text: `加载文件列表失败: ${error.message}`
        });
    });
    */

    downloadForm.addEventListener('submit', function (e) {
        e.preventDefault();

        const fileId = document.getElementById('fileId').value;

        fetch(`http://localhost:4407/download/${fileId}`, {
            method: 'GET',
            headers: {
                'Authorization': token
            }
        })
           .then(response => {
                if (response.ok) {
                    return response.blob();
                }
                return response.json().then(data => { throw new Error(data.error); });
            })
           .then(blob => {
                const url = window.URL.createObjectURL(blob);
                const a = document.createElement('a');
                a.href = url;
                a.download = `file_${fileId}`;
                a.click();
                window.URL.revokeObjectURL(url);
                Swal.fire({
                    icon: 'success',
                    title: '成功',
                    text: '文件下载成功'
                });
            })
           .catch(error => {
                Swal.fire({
                    icon: 'error',
                    title: '错误',
                    text: `下载失败: ${error.message}`
                });
            });
    });
});
