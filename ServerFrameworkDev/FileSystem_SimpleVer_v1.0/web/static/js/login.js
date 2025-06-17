document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('loginForm');

    loginForm.addEventListener('submit', function (e) {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        fetch('http://localhost:4407/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
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
                    localStorage.setItem('token', data.token);
                    Swal.fire({
                        icon: 'success',
                        title: '成功',
                        text: '登录成功'
                    }).then(() => {
                        // 登录成功后跳转到上传页面
                        window.location.href = 'upload.html';
                    });
                }
            })
            .catch(error => {
                Swal.fire({
                    icon: 'error',
                    title: '错误',
                    text: `登录失败: ${error.message}`
                });
            });
    });
});
