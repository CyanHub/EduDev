// 上传文件函数
async function uploadFiles() {
    const fileInput = document.getElementById('fileInput');
    const progressBar = document.getElementById('uploadProgress');

    if (!fileInput.files.length) {
        Swal.fire('提示', '请选择要上传的文件', 'info');
        return;
    }

    progressBar.classList.remove('d-none');
    const progress = progressBar.querySelector('.progress-bar');

    try {
        const formData = new FormData();
        for (let i = 0; i < fileInput.files.length; i++) {
            formData.append('files', fileInput.files[i]);
        }

        const response = await fetch('/api/file/upload', {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            },
            body: formData,
            // 添加上传进度监控
            onUploadProgress: (e) => {
                const percent = Math.round((e.loaded / e.total) * 100);
                progress.style.width = `${percent}%`;
                progress.textContent = `${percent}%`;
            }
        });

        if (response.ok) {
            Swal.fire('成功', '文件上传成功', 'success');
            loadFileList(); // 刷新文件列表
            fileInput.value = ''; // 清空文件选择
            progressBar.classList.add('d-none');
        } else {
            throw new Error('上传失败');
        }
    } catch (error) {
        Swal.fire('错误', error.message, 'error');
        progressBar.classList.add('d-none');
    }
}

// 全局变量存储当前操作的文件ID
let currentFileId = null;
let currentUserId = null;
let isAdmin = false;

// 加载文件列表
async function loadFileList() {
    try {
        const response = await fetch('/api/file/list', {
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });

        const result = await response.json();

        if (response.ok) {
            renderFiles(result.data);
        } else {
            throw new Error(result.msg || '获取文件列表失败');
        }
    } catch (error) {
        console.error('加载文件列表失败:', error);
        Swal.fire('错误', '加载文件列表失败', 'error');
    }
}

// 渲染文件列表
function renderFiles(files) {
    const tbody = document.getElementById('file-list');
    tbody.innerHTML = '';

    if (!files || files.length === 0) {
        tbody.innerHTML = `
            <tr>
                <td colspan="5" class="text-center py-4 text-muted">
                    <i class="bi bi-folder-x" style="font-size: 2rem;"></i>
                    <div class="mt-2">暂无文件</div>
                </td>
            </tr>
        `;
        return;
    }

    files.forEach(file => {
        const canManage = file.uploaderId === currentUserId || isAdmin;
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${file.name}</td>
            <td>${formatFileSize(file.size)}</td>
            <td>${file.uploader || '系统'}</td>
            <td>${new Date(file.createdAt).toLocaleString()}</td>
            <td>
                <button class="btn btn-sm btn-success" 
                        onclick="downloadFile('${file.id}')"
                        ${file.permissions.includes('read') ? '' : 'disabled'}>
                    <i class="bi bi-download"></i> 下载
                </button>
                <button class="btn btn-sm btn-danger" 
                        onclick="deleteFile('${file.id}')"
                        ${canManage ? '' : 'disabled'}>
                    <i class="bi bi-trash"></i> 删除
                </button>
                ${isAdmin ? `
                <button class="btn btn-sm btn-warning" 
                        onclick="showPermissionModal('${file.id}')">
                    <i class="bi bi-shield-lock"></i> 权限
                </button>
                ` : ''}
            </td>
        `;
        tbody.appendChild(tr);
    });
}

// 模拟 formatFileSize 函数
function formatFileSize(bytes) {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

// 模拟 downloadFile 函数
async function downloadFile(fileId) {
    try {
        const response = await fetch(`/api/file/download/${fileId}`, {
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });

        if (response.ok) {
            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = `file_${fileId}`;
            a.click();
            window.URL.revokeObjectURL(url);
        } else {
            throw new Error('下载文件失败');
        }
    } catch (error) {
        console.error('下载文件失败:', error);
        Swal.fire('错误', '下载文件失败', 'error');
    }
}

// 模拟 deleteFile 函数
async function deleteFile(fileId) {
    try {
        const response = await fetch(`/api/file/delete/${fileId}`, {
            method: 'DELETE',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });

        const result = await response.json();

        if (response.ok) {
            Swal.fire({
                icon: 'success',
                title: '删除成功',
                showConfirmButton: false,
                timer: 1500
            });
            loadFileList();
        } else {
            throw new Error(result.msg || '删除文件失败');
        }
    } catch (error) {
        console.error('删除文件失败:', error);
        Swal.fire('错误', '删除文件失败', 'error');
    }
}

// 模拟 showPermissionModal 函数
function showPermissionModal(fileId) {
    currentFileId = fileId;
    const modal = new bootstrap.Modal(document.getElementById('permissionModal'));
    modal.show();
}

// 模拟 setFilePermissions 函数
async function setFilePermissions() {
    if (!currentFileId) return;

    try {
        const targetUserId = document.getElementById('targetUserId').value;
        const permRead = document.getElementById('permRead').checked;
        const permWrite = document.getElementById('permWrite').checked;
        const permDelete = document.getElementById('permDelete').checked;

        const response = await fetch(`/api/file/permissions/${currentFileId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            },
            body: JSON.stringify({
                targetUserId,
                permRead,
                permWrite,
                permDelete
            })
        });

        const result = await response.json();

        if (response.ok) {
            Swal.fire({
                icon: 'success',
                title: '权限设置成功',
                showConfirmButton: false,
                timer: 1500
            });
            const modal = bootstrap.Modal.getInstance(document.getElementById('permissionModal'));
            modal.hide();
            loadFileList();
        } else {
            throw new Error(result.msg || '权限设置失败');
        }
    } catch (error) {
        console.error('权限设置失败:', error);
        Swal.fire('错误', '权限设置失败', 'error');
    }
}

// 模拟 uploadFiles 函数
async function uploadFiles() {
    const fileInput = document.getElementById('fileInput');
    const files = fileInput.files;
    if (files.length === 0) {
        Swal.fire('提示', '请选择要上传的文件', 'info');
        return;
    }

    const formData = new FormData();
    for (let i = 0; i < files.length; i++) {
        formData.append('files', files[i]);
    }

    try {
        const response = await fetch('/api/file/upload', {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            },
            body: formData
        });

        const result = await response.json();

        if (response.ok) {
            await Swal.fire({
                icon: 'success',
                title: '上传成功',
                showConfirmButton: false,
                timer: 1500
            });
            loadFileList();
            fileInput.value = ''; // 清空文件选择
        } else {
            throw new Error(result.msg || '上传失败');
        }
    } catch (error) {
        Swal.fire('上传失败', error.message, 'error');
    }
}

// 模拟 logout 函数
async function logout() {
    try {
        const response = await fetch('/api/auth/logout', {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });

        if (response.ok) {
            localStorage.removeItem('token');
            window.location.href = 'login.html';
        } else {
            throw new Error('登出失败');
        }
    } catch (error) {
        console.error('登出失败:', error);
        Swal.fire('错误', '登出失败', 'error');
    }
}

// 页面加载时加载文件列表
document.addEventListener('DOMContentLoaded', () => {
    loadFileList();
    // 可添加获取 currentUserId 和 isAdmin 的逻辑
});