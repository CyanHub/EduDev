import React, { createContext, useState, useContext, useEffect } from 'react';
import { message } from 'antd';
import { jwtDecode } from 'jwt-decode';
import api from '../utils/api';

const AuthContext = createContext(null);

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [loading, setLoading] = useState(true);

  // 初始化时从本地存储加载用户信息
  useEffect(() => {
    const initAuth = () => {
      const token = localStorage.getItem('token');
      const userInfo = localStorage.getItem('user');
      
      if (token && userInfo) {
        try {
          const decodedToken = jwtDecode(token);
          const currentTime = Date.now() / 1000;
          
          if (decodedToken.exp > currentTime) {
            setUser(JSON.parse(userInfo));
            setIsAuthenticated(true);
            api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
          } else {
            // 令牌已过期，清除本地存储
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            delete api.defaults.headers.common['Authorization'];
          }
        } catch (error) {
          console.error('令牌解析错误:', error);
          localStorage.removeItem('token');
          localStorage.removeItem('user');
        }
      }
      setLoading(false);
    };

    initAuth();
  }, []);

  // 用户登录
  const login = async (credentials, isAdmin = false) => {
    try {
      setLoading(true);
      const endpoint = isAdmin ? '/admin/login' : '/users/login';
      const response = await api.post(endpoint, credentials);
      
      if (response.data.code === 200) {
        const { token, ...userData } = response.data.data;
        
        // 保存令牌和用户信息到本地存储
        localStorage.setItem('token', token);
        localStorage.setItem('user', JSON.stringify(userData));
        
        // 设置API请求头的认证令牌
        api.defaults.headers.common['Authorization'] = `Bearer ${token}`;
        
        setUser(userData);
        setIsAuthenticated(true);
        message.success('登录成功');
        return true;
      } else {
        message.error(response.data.message || '登录失败');
        return false;
      }
    } catch (error) {
      console.error('登录错误:', error);
      message.error(error.response?.data?.message || '登录失败，请稍后再试');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 用户注册
  const register = async (userData) => {
    try {
      setLoading(true);
      const response = await api.post('/users/register', userData);
      
      if (response.data.code === 200) {
        message.success('注册成功，请登录');
        return true;
      } else {
        message.error(response.data.message || '注册失败');
        return false;
      }
    } catch (error) {
      console.error('注册错误:', error);
      message.error(error.response?.data?.message || '注册失败，请稍后再试');
      return false;
    } finally {
      setLoading(false);
    }
  };

  // 用户登出
  const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    delete api.defaults.headers.common['Authorization'];
    setUser(null);
    setIsAuthenticated(false);
    message.success('已退出登录');
  };

  // 更新用户信息
  const updateUserInfo = async () => {
    if (!isAuthenticated || !user) return;
    
    try {
      const response = await api.get(`/users/${user.user_id}`);
      
      if (response.data.code === 200) {
        const updatedUser = response.data.data;
        setUser(updatedUser);
        localStorage.setItem('user', JSON.stringify(updatedUser));
      }
    } catch (error) {
      console.error('获取用户信息错误:', error);
    }
  };

  const value = {
    user,
    isAuthenticated,
    loading,
    login,
    register,
    logout,
    updateUserInfo
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};