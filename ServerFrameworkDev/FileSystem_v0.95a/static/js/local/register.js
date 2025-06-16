// 头像预览功能保持不变
const avatarPreview = document.getElementById('avatarPreview');
const avatarInput = document.getElementById('avatarInput');

avatarPreview.addEventListener('click', () => {
    avatarInput.click();
});

avatarInput.addEventListener('change', (e) => {
    const file = e.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (event) => {
            avatarPreview.innerHTML = `<img src="${event.target.result}" alt="头像预览">`;
        };
        reader.readAsDataURL(file);
    }
});

// 注册表单提交
document.getElementById('register-form').addEventListener('submit', async (e) => {
    e.preventDefault();

    const username = document.getElementById('username').value.trim();
    const password = document.getElementById('password').value.trim();
    const confirmPassword = document.getElementById('confirm-password').value.trim();
    const email = document.getElementById('email').value.trim();
    const roleId = document.getElementById('role').value;

    if (roleId === '0') {
        Swal.fire({
            icon: 'error',
            title: '角色未选择',
            text: '请选择一个角色',
            showConfirmButton: true
        });
        return;
    }

    if (!username || !password || !confirmPassword || !email || !roleId) {
        Swal.fire({
            icon: 'error',
            title: '信息不完整',
            text: '请填写所有必填字段',
            showConfirmButton: true
        });
        return;
    }

    if (password !== confirmPassword) {
        Swal.fire({
            icon: 'error',
            title: '密码不一致',
            text: '两次输入的密码不匹配',
            showConfirmButton: true
        });
        return;
    }

    try {
        const formData = new FormData();
        formData.append('username', username);
        formData.append('password', password);
        formData.append('nickname', document.getElementById('nickname').value);
        formData.append('email', email);
        formData.append('phone', document.getElementById('phone').value);
        formData.append('roleId', roleId);

        // 增加空值检查
        const avatarInput = document.getElementById('avatarInput');
        if (avatarInput && avatarInput.files[0]) {
            formData.append('avatar', avatarInput.files[0]);
        }

        const response = await fetch('/api/user/register', {
            method: 'POST',
            body: formData
        });

        const result = await response.json();

        if (response.ok) {
            Swal.fire({
                icon: 'success',
                title: '注册成功',
                text: '请登录',
                showConfirmButton: true
            }).then(() => {
                window.location.href = 'login.html';
            });
        } else {
            throw new Error(result.message || '注册失败');
        }
    } catch (error) {
        console.error('注册失败:', error);
        Swal.fire('错误', error.message, 'error');
    }
});
