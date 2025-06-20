document.addEventListener('DOMContentLoaded', function () {
    // 初始化提示框
    const tooltipTriggerList = document.querySelectorAll('[data-bs-toggle="tooltip"]')
    const tooltipList = [...tooltipTriggerList].map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl))

    const uploadForm = document.getElementById('uploadForm');
    const token = localStorage.getItem('token');

    if (!token) {
        Swal.fire({
            icon: 'error',
            title: '错误',
            text: '请先登录'
        });
        return;
    }

    uploadForm.addEventListener('submit', async function (e) {
        e.preventDefault();

        const fileInput = document.getElementById('file');
        const file = fileInput.files[0];
        const isPublic = document.getElementById('isPublic').checked;

        // 检查文件名是否重复
        const newFileName = await checkFileName(file.name, token);

        const formData = new FormData();
        formData.append('file', new File([file], newFileName, { type: file.type }));
        formData.append('isPublic', isPublic);

        fetch('http://localhost:4407/upload', {
            method: 'POST',
            headers: {
                'Authorization': token
            },
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    Swal.fire({
                        icon: 'error',
                        title: '错误',
                        text: data.error
                    });
                } else {
                    Swal.fire({
                        icon: 'success',
                        title: '成功',
                        html: `文件上传成功 <br> 文件ID为: ${data.fileId} <br> 您可以使用文件ID来下载文件`
                    });
                }
            })
            .catch(error => {
                Swal.fire({
                    icon: 'error',
                    title: '错误',
                    text: `上传失败: ${error.message}`
                });
            });
    });

    async function checkFileName(originalName, token) {
        let newName = originalName;
        let counter = 1;
        const extIndex = originalName.lastIndexOf('.');
        const nameWithoutExt = extIndex === -1 ? originalName : originalName.slice(0, extIndex);
        const ext = extIndex === -1 ? '' : originalName.slice(extIndex);

        while (true) {
            const response = await fetch(`http://localhost:4407/check-file-name?name=${encodeURIComponent(newName)}`, {
                method: 'GET',
                headers: {
                    'Authorization': token
                }
            });
            const data = await response.json();
            if (!data.exists) {
                if (counter > 1) {
                    Swal.fire({
                        icon: 'info',
                        title: '提示',
                        text: `有同名文件，已改名为 ${newName}`
                    });
                }
                break;
            }
            newName = `${nameWithoutExt}${String(counter).padStart(2, '0')}${ext}`;
            counter++;
        }
        return newName;
    }
});
