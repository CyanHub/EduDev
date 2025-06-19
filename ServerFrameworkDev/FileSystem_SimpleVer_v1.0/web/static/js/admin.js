document.addEventListener('DOMContentLoaded', function () {
    const token = localStorage.getItem('token');
    if (!token) {
        Swal.fire({
            icon: 'error',
            title: '错误',
            text: '请先登录'
        }).then(() => {
            window.location.href = 'login.html';
        });
        return;
    }

    loadUsers();
    loadRoles();
    loadPermissions();
});

function loadUsers() {
    const token = localStorage.getItem('token');
    fetch('http://localhost:4407/users', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
        .then(response => response.json())
        .then(data => {
            const usersList = document.getElementById('usersList');
            usersList.innerHTML = '';
            data.forEach(user => {
                const userDiv = document.createElement('div');
                userDiv.className = 'mb-2';
                userDiv.innerHTML = `
                    <div class="d-flex justify-content-between align-items-center">
                        <span>${user.Username}</span>
                        <div>
                            <button class="btn btn-sm btn-primary" onclick="assignRoleToUser(${user.ID})">分配角色</button>
                            <button class="btn btn-sm btn-danger" onclick="deleteUser(${user.ID})">删除</button>
                        </div>
                    </div>
                `;
                usersList.appendChild(userDiv);
            });
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `加载用户列表失败: ${error.message}`
            });
        });
}

function loadRoles() {
    const token = localStorage.getItem('token');
    fetch('http://localhost:4407/roles', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
        .then(response => response.json())
        .then(data => {
            const rolesList = document.getElementById('rolesList');
            rolesList.innerHTML = '';
            data.forEach(role => {
                const roleDiv = document.createElement('div');
                roleDiv.className = 'mb-2';
                roleDiv.innerHTML = `
                    <div class="d-flex justify-content-between align-items-center">
                        <span>${role.Name}</span>
                        <div>
                            <button class="btn btn-sm btn-primary" onclick="assignPermissionToRole(${role.ID})">分配权限</button>
                            <button class="btn btn-sm btn-danger" onclick="deleteRole(${role.ID})">删除</button>
                        </div>
                    </div>
                `;
                rolesList.appendChild(roleDiv);
            });
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `加载角色列表失败: ${error.message}`
            });
        });
}

function loadPermissions() {
    const token = localStorage.getItem('token');
    fetch('http://localhost:4407/permissions', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
        .then(response => response.json())
        .then(data => {
            const permissionsList = document.getElementById('permissionsList');
            permissionsList.innerHTML = '';
            data.forEach(permission => {
                const permissionDiv = document.createElement('div');
                permissionDiv.className = 'mb-2';
                permissionDiv.innerHTML = `
                    <div class="d-flex justify-content-between align-items-center">
                        <span>${permission.Name} (${permission.Code})</span>
                        <button class="btn btn-sm btn-danger" onclick="deletePermission(${permission.ID})">删除</button>
                    </div>
                `;
                permissionsList.appendChild(permissionDiv);
            });
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `加载权限列表失败: ${error.message}`
            });
        });
}

function addUser() {
    // 处理第一个添加用户按钮的逻辑
    const username = document.getElementById('userUsername').value;
    const password = document.getElementById('userPassword').value;
    const email = document.getElementById('userEmail').value;

    fetch('http://localhost:4407/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({
            username: username,
            password: password,
            email: email
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
                    text: '用户添加成功'
                }).then(() => {
                    loadUsers();
                    $('#userModal').modal('hide');
                });
            }
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `添加用户失败: ${error.message}`
            });
        });
}

// function addUser2() {
//     // 处理第二个添加用户按钮的逻辑
//     const username = document.getElementById('userUsername2').value;
//     const password = document.getElementById('userPassword2').value;
//     const email = document.getElementById('userEmail2').value;

//     fetch('http://localhost:4407/register', {
//         method: 'POST',
//         headers: {
//             'Content-Type': 'application/json',
//             'Authorization': token
//         },
//         body: JSON.stringify({
//             username: username,
//             password: password,
//             email: email
//         })
//     })
//         .then(response => response.json())
//         .then(data => {
//             if (data.error) {
//                 Swal.fire({
//                     icon: 'error',
//                     title: '错误',
//                     text: data.error
//                 });
//             } else {
//                 Swal.fire({
//                     icon: 'success',
//                     title: '成功',
//                     text: '用户添加成功'
//                 }).then(() => {
//                     loadUsers();
//                     $('#userModal').modal('hide');
//                 });
//             }
//         })
//         .catch(error => {
//             Swal.fire({
//                 icon: 'error',
//                 title: '错误',
//                 text: `添加用户失败: ${error.message}`
//             });
//         });
// }

function addRole() {
    const token = localStorage.getItem('token');
    const roleName = document.getElementById('roleName').value;

    fetch('http://localhost:4407/roles', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({
            name: roleName
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
                    text: '角色添加成功'
                }).then(() => {
                    loadRoles();
                    $('#roleModal').modal('hide');
                });
            }
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `添加角色失败: ${error.message}`
            });
        });
}

function addPermission() {
    const token = localStorage.getItem('token');
    const permissionName = document.getElementById('permissionName').value;
    const permissionCode = document.getElementById('permissionCode').value;

    fetch('http://localhost:4407/permissions', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({
            name: permissionName,
            code: permissionCode
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
                    text: '权限添加成功'
                }).then(() => {
                    loadPermissions();
                    $('#permissionModal').modal('hide');
                });
            }
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `添加权限失败: ${error.message}`
            });
        });
}

function assignRoleToUser(userId) {
    const token = localStorage.getItem('token');
    fetch('http://localhost:4407/roles', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log('Fetched roles:', data);
            // 正确构建角色选项，value 为角色 ID，label 为角色名称
            const roles = data.reduce((acc, role) => {
                acc[role.ID] = role.Name;
                return acc;
            }, {});

            Swal.fire({
                title: '分配角色',
                input: 'select',
                inputOptions: roles,
                inputPlaceholder: '选择角色',
                showCancelButton: true,
                inputValidator: (value) => {
                    if (!value) {
                        return '请选择一个角色';
                    }
                    return null;
                }
            }).then((result) => {
                if (result.isConfirmed) {
                    const roleId = Number(result.value);
                    console.log('Selected role ID value:', result.value);
                    console.log('Parsed role ID:', roleId);
                    if (isNaN(roleId)) {
                        Swal.fire({
                            icon: 'error',
                            title: '错误',
                            text: '选择的角色 ID 无效'
                        });
                        return;
                    }
                    fetch('http://localhost:4407/user-roles', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': token
                        },
                        body: JSON.stringify({
                            user_id: userId,
                            role_id: roleId
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
                                    text: '角色分配成功'
                                });
                            }
                        })
                        .catch(error => {
                            Swal.fire({
                                icon: 'error',
                                title: '错误',
                                text: `分配角色失败: ${error.message}`
                            });
                        });
                }
            });
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `加载角色列表失败: ${error.message}`
            });
        });
}

function assignPermissionToRole(roleId) {
    const token = localStorage.getItem('token');
    fetch('http://localhost:4407/permissions', {
        method: 'GET',
        headers: {
            'Authorization': token
        }
    })
        .then(response => response.json())
        .then(data => {
            const permissions = data.map(permission => ({
                value: permission.ID,
                label: `${permission.Name} (${permission.Code})`
            }));

            Swal.fire({
                title: '分配权限',
                input: 'select',
                inputOptions: permissions,
                inputPlaceholder: '选择权限',
                showCancelButton: true,
                inputValidator: (value) => {
                    if (!value) {
                        return '请选择一个权限';
                    }
                    return null;
                }
            }).then((result) => {
                if (result.isConfirmed) {
                    const permissionId = result.value;
                    fetch('http://localhost:4407/role-permissions', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                            'Authorization': token
                        },
                        body: JSON.stringify({
                            role_id: roleId,
                            permission_id: permissionId
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
                                    text: '权限分配成功'
                                });
                            }
                        })
                        .catch(error => {
                            Swal.fire({
                                icon: 'error',
                                title: '错误',
                                text: `分配权限失败: ${error.message}`
                            });
                        });
                }
            });
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `加载权限列表失败: ${error.message}`
            });
        });
}

function deleteUser(userId) {
    const token = localStorage.getItem('token');
    Swal.fire({
        title: '确定要删除该用户吗？',
        text: '删除后将再也无法恢复，用户名在再也无法使用！！！',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: '删除'
    }).then((result) => {
        if (result.isConfirmed) {
            fetch(`http://localhost:4407/users/${userId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': token
                }
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
                            text: '用户删除成功'
                        }).then(() => {
                            loadUsers();
                        });
                    }
                })
                .catch(error => {
                    Swal.fire({
                        icon: 'error',
                        title: '错误',
                        text: `删除用户失败: ${error.message}`
                    });
                });
        }
    });
}

function deleteRole(roleId) {
    const token = localStorage.getItem('token');
    Swal.fire({
        title: '确定要删除该角色吗？',
        text: '删除后将再也无法恢复，角色名再也无法使用！！！',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: '删除'
    }).then((result) => {
        if (result.isConfirmed) {
            fetch(`http://localhost:4407/roles/${roleId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': token
                }
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
                            text: '角色删除成功'
                        }).then(() => {
                            loadRoles();
                        });
                    }
                })
                .catch(error => {
                    Swal.fire({
                        icon: 'error',
                        title: '错误',
                        text: `删除角色失败: ${error.message}`
                    });
                });
        }
    });
}

function deletePermission(permissionId) {
    const token = localStorage.getItem('token');
    Swal.fire({
        title: '确定要删除该权限吗？',
        text: '删除后将再也无法恢复，权限名再也无法使用！！！',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#d33',
        cancelButtonColor: '#3085d6',
        confirmButtonText: '删除'
    }).then((result) => {
        if (result.isConfirmed) {
            fetch(`http://localhost:4407/permissions/${permissionId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': token
                }
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
                            text: '权限删除成功'
                        }).then(() => {
                            loadPermissions();
                        });
                    }
                })
                .catch(error => {
                    Swal.fire({
                        icon: 'error',
                        title: '错误',
                        text: `删除权限失败: ${error.message}`
                    });
                });
        }
    });
}


function addAdmin() {
    const token = localStorage.getItem('token');
    const username = document.getElementById('adminUsername').value;
    const password = document.getElementById('adminPassword').value;
    const email = document.getElementById('adminEmail').value;

    fetch('http://localhost:4407/add-admin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({
            username: username,
            password: password,
            email: email
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
                    text: '管理者添加成功'
                }).then(() => {
                    loadUsers();
                    $('#addAdminModal').modal('hide');
                });
            }
        })
        .catch(error => {
            Swal.fire({
                icon: 'error',
                title: '错误',
                text: `添加管理者失败: ${error.message}`
            });
        });

}
// 模态框关闭事件，添加对 DOM 元素是否存在的检查
const element = document.getElementById('your-element-id');
if (element) {
    element.addEventListener('event-name', function () {
        // 事件处理逻辑
        // 假设这是模态框关闭按钮的点击事件处理函数
        document.getElementById('closeModalButton').addEventListener('click', function () {
            const modalElement = document.getElementById('userModal');
            const modal = bootstrap.Modal.getInstance(modalElement);
            if (modal) {
                modal.hide();
            }
        });
    });
}
