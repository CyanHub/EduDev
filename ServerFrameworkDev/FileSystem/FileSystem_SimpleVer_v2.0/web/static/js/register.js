document.addEventListener('DOMContentLoaded', function () {
    // 初始化提示框
    const tooltipTriggerList = document.querySelectorAll('[data-bs-toggle="tooltip"]')
    const tooltipList = [...tooltipTriggerList].map(tooltipTriggerEl => new bootstrap.Tooltip(tooltipTriggerEl))

    const registerForm = document.getElementById('registerForm');

    registerForm.addEventListener('submit', function (e) {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        // 验证 Username 为字母数字组合
        const usernameRegex = /^[a-zA-Z0-9]+$/;
        if (!usernameRegex.test(username)) {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: '用户名必须为字母数字组合'
            });
            return;
        }
        // 验证是否超过或者小于规定的字符串
        if (username.length < 4 || username.length > 16) {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: '用户名长度必须在 4 到 16 个字符之间'
            });
            return;
        }

        // 验证 Password 包含特定字符，假设至少包含一个大写字母、一个小写字母和一个数字
        const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/;
        if (!passwordRegex.test(password)) {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: '密码必须包含至少一个大写字母、一个小写字母和一个数字'
            });
            return;
        }

        fetch('http://localhost:4407/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: password,
                email: document.getElementById('email').value
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
                    Swal.fire({
                        icon: 'success',
                        title: '成功',
                        text: data.message
                    }).then(() => {
                        // 注册成功后跳转到登录页面
                        window.location.href = 'login.html';
                    });
                }
            })
            .catch(error => {
                Swal.fire({
                    icon: 'error',
                    title: '错误',
                    text: `注册失败: ${error.message}`
                });
            });
    });
});
