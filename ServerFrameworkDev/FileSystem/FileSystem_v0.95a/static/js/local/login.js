document.addEventListener('DOMContentLoaded', () => {
    const usernameInput = document.getElementById('username');
    const avatarContainer = document.getElementById('avatar-container');
    const userAvatar = document.getElementById('user-avatar');
    const inputPrompt = avatarContainer.querySelector('p');

    // 防抖函数
    function debounce(func, delay) {
        let timer = null;
        return function() {
            const context = this;
            const args = arguments;
            clearTimeout(timer);
            timer = setTimeout(() => {
                func.apply(context, args);
            }, delay);
        };
    }

    // 使用防抖函数包装头像检测逻辑
    const checkAvatar = debounce(async () => {
        const username = usernameInput.value.trim();
        if (username) {
            try {
                const response = await fetch(`/api/user/avatar?username=${username}`);
                const data = await response.json();
                if (response.ok && data.avatarPath) {
                    inputPrompt.style.display = 'none';
                    userAvatar.src = data.avatarPath;
                    userAvatar.style.display = 'block';
                } else {
                    inputPrompt.style.display = 'block';
                    userAvatar.style.display = 'none';
                }
            } catch (error) {
                console.error('获取头像信息失败:', error);
            }
        } else {
            inputPrompt.style.display = 'block';
            userAvatar.style.display = 'none';
        }
    }, 3000);

    usernameInput.addEventListener('input', checkAvatar);

    const loginForm = document.getElementById('login-form');
    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const loginButton = document.querySelector('#login-form button[type="submit"]');
            const originalButtonText = loginButton.textContent;

            try {
                // 显示加载状态
                loginButton.disabled = true;
                loginButton.textContent = '登录中...';

                const username = document.getElementById('username').value.trim();
                const password = document.getElementById('password').value.trim();

                if (!username || !password) {
                    Swal.fire({
                        icon: 'error',
                        title: '信息不完整',
                        text: '请填写用户名和密码',
                        showConfirmButton: true
                    });
                    return;
                }

                const response = await fetch('/api/user/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        username: username,
                        password: password
                    })
                });

                const result = await response.json();

                if (response.ok) {
                    // 存储 token 和用户 ID
                    if (result.data.token) {
                        localStorage.setItem('token', result.data.token);
                    }
                    if (result.data.id) {
                        localStorage.setItem('user_id', result.data.id);
                    }

                    // 获取用户信息
                    if (result.data.token) {
                        const userInfoResponse = await fetch('/api/user/info', {
                            headers: {
                                'Authorization': `Bearer ${result.data.token}`
                            }
                        });
                        const userInfo = await userInfoResponse.json();
                        if (userInfoResponse.ok) {
                            // 存储用户信息
                            localStorage.setItem('user_name', userInfo.data.username);
                            localStorage.setItem('user_email', userInfo.data.email);
                            localStorage.setItem('user_role', userInfo.data.role);
                        } else {
                            console.error('获取用户信息失败:', userInfo.msg);
                        }
                    }

                    await Swal.fire({
                        icon: 'success',
                        title: '登录成功',
                        text: '正在跳转到文件管理页面...',
                        showConfirmButton: false,
                        timer: 1500
                    });
                    window.location.href = 'list.html';
                } else {
                    throw new Error(result.message || '登录失败');
                }
            } catch (error) {
                console.error('登录失败:', error);
                Swal.fire({
                    icon: 'error',
                    title: '登录失败',
                    text: error.message,
                    showConfirmButton: true
                });
            } finally {
                // 恢复按钮状态
                loginButton.disabled = false;
                loginButton.textContent = originalButtonText;
            }
        });
    }
});

