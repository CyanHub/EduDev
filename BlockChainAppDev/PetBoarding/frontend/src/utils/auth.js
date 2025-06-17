import { jwtDecode } from 'jwt-decode';

/**
 * 检查JWT令牌是否过期
 * @returns {boolean} 如果令牌过期或无效则返回true，否则返回false
 */
export const checkTokenExpiration = () => {
  const token = localStorage.getItem('token');
  
  if (!token) {
    return true; // 没有令牌，视为已过期
  }
  
  try {
    const decodedToken = jwtDecode(token);
    const currentTime = Date.now() / 1000; // 转换为秒
    
    if (decodedToken.exp < currentTime) {
      // 令牌已过期
      return true;
    }
    
    // 令牌有效
    return false;
  } catch (error) {
    console.error('令牌解析错误:', error);
    return true; // 解析错误，视为已过期
  }
};

/**
 * 获取当前用户角色
 * @returns {string|null} 用户角色或null（如果未登录）
 */
export const getUserRole = () => {
  const userInfo = localStorage.getItem('user');
  
  if (!userInfo) {
    return null;
  }
  
  try {
    const user = JSON.parse(userInfo);
    return user.role;
  } catch (error) {
    console.error('用户信息解析错误:', error);
    return null;
  }
};

/**
 * 检查用户是否有权限访问特定角色的路由
 * @param {string} requiredRole 所需角色
 * @returns {boolean} 如果用户有权限则返回true，否则返回false
 */
export const hasPermission = (requiredRole) => {
  if (!requiredRole) {
    return true; // 没有指定所需角色，任何人都可以访问
  }
  
  const userRole = getUserRole();
  
  if (!userRole) {
    return false; // 未登录用户没有权限
  }
  
  // 管理员可以访问任何角色的路由
  if (userRole === 'admin') {
    return true;
  }
  
  // 检查用户角色是否匹配所需角色
  return userRole === requiredRole;
};

/**
 * 从本地存储获取用户ID
 * @returns {string|null} 用户ID或null（如果未登录）
 */
export const getUserId = () => {
  const userInfo = localStorage.getItem('user');
  
  if (!userInfo) {
    return null;
  }
  
  try {
    const user = JSON.parse(userInfo);
    return user.user_id || user.admin_id;
  } catch (error) {
    console.error('用户信息解析错误:', error);
    return null;
  }
};