document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('loginForm');

    loginForm.addEventListener('submit', function (e) {
        e.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        const isAdmin = document.getElementById('isAdmin').checked;

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
                    if (isAdmin) {
                        // 验证用户是否有管理者权限
                        fetch('http://localhost:4407/check-admin', {
                            method: 'GET',
                            headers: {
                                'Authorization': data.token
                            }
                        })
                           .then(response => response.json())
                           .then(adminData => {
                                if (adminData.isAdmin) {
                                    Swal.fire({
                                        icon: 'success',
                                        title: '成功',
                                        text: '管理者登录成功'
                                    }).then(() => {
                                        // 跳转到管理页面
                                        window.location.href = 'admin.html';
                                    });
                                } else {
                                    Swal.fire({
                                        icon: 'error',
                                        title: '错误',
                                        text: '你没有管理者权限'
                                    });
                                }
                            })
                           .catch(error => {
                                Swal.fire({
                                    icon: 'error',
                                    title: '错误',
                                    text: `验证管理者权限失败: ${error.message}`
                                });
                            });
                    } else {
                        Swal.fire({
                            icon: 'success',
                            title: '成功',
                            text: '登录成功'
                        }).then(() => {
                            // 普通用户登录成功后跳转到上传页面
                            window.location.href = 'upload.html';
                        });
                    }
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
