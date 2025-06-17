document.addEventListener('DOMContentLoaded', function () {
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

    uploadForm.addEventListener('submit', function (e) {
        e.preventDefault();

        const fileInput = document.getElementById('file');
        const file = fileInput.files[0];
        const formData = new FormData();
        formData.append('file', file);

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
                        text: data.message
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
});
