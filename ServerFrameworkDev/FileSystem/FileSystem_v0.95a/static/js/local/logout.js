document.addEventListener('DOMContentLoaded', () => {
    const logoutProgress = document.getElementById('logout-progress');
    let progress = 0;
    const interval = setInterval(() => {
        progress += 10;
        logoutProgress.style.width = `${progress}%`;

        if (progress >= 100) {
            clearInterval(interval);
            // 清除所有本地存储的登录信息
            localStorage.removeItem('token');
            localStorage.removeItem('user_id');
            localStorage.removeItem('user_name');
            localStorage.removeItem('user_email');
            localStorage.removeItem('user_role');
            // 可以根据实际情况添加需要清除的其他用户信息

            window.location.href = 'login.html';
        }
    }, 200);
});