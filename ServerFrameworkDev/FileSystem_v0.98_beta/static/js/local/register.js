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
    const roleName = document.getElementById('role').value;

    // 将角色名称转换为角色 ID
    let roleId;
    if (roleName === 'admin') {
        roleId = 1;
    } else if (roleName === 'user') {
        roleId = 2;
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
        // 构建JSON请求体
        const requestBody = {
            username,
            password,
            nickname: document.getElementById('nickname').value,
            email,
            phone: document.getElementById('phone').value,
            roleId // 使用角色 ID
        };

        const response = await fetch('/api/user/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestBody)
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