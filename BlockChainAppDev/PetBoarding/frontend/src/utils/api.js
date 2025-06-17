import axios from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: '/api', // 使用相对路径，结合package.json中的proxy配置
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 从本地存储获取令牌
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response) {
      // 服务器返回错误状态码
      const { status, data } = error.response;
      
      switch (status) {
        case 401: // 未授权
          // 如果是令牌过期，清除本地存储并重定向到登录页面
          if (window.location.pathname !== '/login' && window.location.pathname !== '/admin/login') {
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            window.location.href = '/login';
          }
          break;
        case 403: // 禁止访问
          console.error('禁止访问:', data.message);
          break;
        case 404: // 资源不存在
          console.error('资源不存在:', data.message);
          break;
        case 500: // 服务器错误
          console.error('服务器错误:', data.message);
          break;
        default:
          console.error(`错误 ${status}:`, data.message);
      }
    } else if (error.request) {
      // 请求已发送但未收到响应
      console.error('网络错误，未收到响应:', error.request);
    } else {
      // 请求配置出错
      console.error('请求错误:', error.message);
    }
    
    return Promise.reject(error);
  }
);

export default api;

// 用户相关API
export const userApi = {
  register: (data) => api.post('/users/register', data),
  login: (data) => api.post('/users/login', data),
  getInfo: (userId) => api.get(`/users/${userId}`),
  updateInfo: (userId, data) => api.put(`/users/${userId}`, data),
};

// 宠物相关API
export const petApi = {
  getAll: (userId) => api.get(`/pets/user/${userId}`),
  getOne: (petId) => api.get(`/pets/${petId}`),
  create: (data) => api.post('/pets', data),
  update: (petId, data) => api.put(`/pets/${petId}`, data),
  delete: (petId) => api.delete(`/pets/${petId}`),
};

// 寄养服务相关API
export const serviceApi = {
  search: (params) => api.get('/boarding/search', { params }),
  getOne: (serviceId) => api.get(`/boarding/${serviceId}`),
  create: (data) => api.post('/boarding', data),
  update: (serviceId, data) => api.put(`/boarding/${serviceId}`, data),
};

// 订单相关API
export const orderApi = {
  getUserOrders: (userId, params) => api.get(`/orders/user/${userId}`, { params }),
  getProviderOrders: (providerId, params) => api.get(`/orders/provider/${providerId}`, { params }),
  getOne: (orderId) => api.get(`/orders/${orderId}`),
  create: (data) => api.post('/orders', data),
  updateStatus: (orderId, data) => api.put(`/orders/${orderId}/status`, data),
  cancel: (orderId, data) => api.put(`/orders/${orderId}/cancel`, data),
};

// 评价相关API
export const reviewApi = {
  getServiceReviews: (serviceId, params) => api.get(`/reviews/service/${serviceId}`, { params }),
  getUserReviews: (userId, params) => api.get(`/reviews/user/${userId}`, { params }),
  getOne: (reviewId) => api.get(`/reviews/${reviewId}`),
  create: (data) => api.post('/reviews', data),
};

// 通知相关API
export const notificationApi = {
  getUserNotifications: (userId, params) => api.get(`/notifications/user/${userId}`, { params }),
  getOne: (notificationId) => api.get(`/notifications/${notificationId}`),
  markAsRead: (notificationId) => api.put(`/notifications/${notificationId}/read`),
  markAllAsRead: (userId) => api.put(`/notifications/user/${userId}/read-all`),
};

// 管理员相关API
export const adminApi = {
  login: (data) => api.post('/admin/login', data),
  getDashboard: () => api.get('/admin/dashboard'),
  getUsers: (params) => api.get('/admin/users', { params }),
  updateUserStatus: (userId, data) => api.put(`/admin/users/${userId}/status`, data),
  getServices: (params) => api.get('/admin/services', { params }),
  updateServiceStatus: (serviceId, data) => api.put(`/admin/services/${serviceId}/status`, data),
};

// 健康检查API
export const healthApi = {
  check: () => api.get('/health'),
};