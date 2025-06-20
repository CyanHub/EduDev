// 检查是否已登录
document.addEventListener('DOMContentLoaded', () => {
    const token = localStorage.getItem('token');
    if (token) {
        Swal.fire({
            title: '检测到登录状态',
            text: '您已登录，是否跳转到文件管理页面？',
            icon: 'info',
            showCancelButton: true,
            confirmButtonText: '前往',
            cancelButtonText: '留在首页',
            focusCancel: true
        }).then((result) => {
            if (result.isConfirmed) {
                window.location.href = 'list.html';
            }
        });
    }
});