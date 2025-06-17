import React, { useState, useEffect } from 'react';
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom';
import { Layout, ConfigProvider, message } from 'antd';
import zhCN from 'antd/locale/zh_CN';

// 组件导入
import Login from './pages/Login';
import Register from './pages/Register';
import Home from './pages/Home';
import UserDashboard from './pages/user/Dashboard';
import PetList from './pages/user/PetList';
import PetDetail from './pages/user/PetDetail';
import ServiceList from './pages/service/ServiceList';
import ServiceDetail from './pages/service/ServiceDetail';
import OrderList from './pages/order/OrderList';
import OrderDetail from './pages/order/OrderDetail';
import ReviewList from './pages/review/ReviewList';
import NotificationList from './pages/notification/NotificationList';
import AdminLogin from './pages/admin/Login';
import AdminDashboard from './pages/admin/Dashboard';
import AdminUserList from './pages/admin/UserList';
import AdminServiceList from './pages/admin/ServiceList';
import AdminOrderList from './pages/admin/OrderList';
import NotFound from './pages/NotFound';

// 工具和上下文
import { AuthProvider, useAuth } from './contexts/AuthContext';
import { checkTokenExpiration } from './utils/auth';
import './App.css';

const { Content } = Layout;

// 受保护的路由组件
const ProtectedRoute = ({ children, requiredRole }) => {
  const { isAuthenticated, user } = useAuth();
  const navigate = useNavigate();

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }

  if (requiredRole && user.role !== requiredRole) {
    message.error('您没有权限访问此页面');
    return <Navigate to="/" replace />;
  }

  return children;
};

// 管理员路由组件
const AdminRoute = ({ children }) => {
  const { isAuthenticated, user } = useAuth();
  const navigate = useNavigate();

  if (!isAuthenticated || user.role !== 'admin') {
    return <Navigate to="/admin/login" replace />;
  }

  return children;
};

function AppContent() {
  const { isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  // 检查令牌是否过期
  useEffect(() => {
    if (isAuthenticated) {
      const checkToken = () => {
        if (checkTokenExpiration()) {
          message.warning('登录已过期，请重新登录');
          logout();
          navigate('/login');
        }
      };

      checkToken();
      const interval = setInterval(checkToken, 60000); // 每分钟检查一次

      return () => clearInterval(interval);
    }
  }, [isAuthenticated, logout, navigate]);

  return (
    <Layout className="app-layout">
      <Content className="app-content">
        <Routes>
          {/* 公共路由 */}
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/services" element={<ServiceList />} />
          <Route path="/services/:id" element={<ServiceDetail />} />

          {/* 用户路由 */}
          <Route path="/dashboard" element={
            <ProtectedRoute>
              <UserDashboard />
            </ProtectedRoute>
          } />
          <Route path="/pets" element={
            <ProtectedRoute>
              <PetList />
            </ProtectedRoute>
          } />
          <Route path="/pets/:id" element={
            <ProtectedRoute>
              <PetDetail />
            </ProtectedRoute>
          } />
          <Route path="/orders" element={
            <ProtectedRoute>
              <OrderList />
            </ProtectedRoute>
          } />
          <Route path="/orders/:id" element={
            <ProtectedRoute>
              <OrderDetail />
            </ProtectedRoute>
          } />
          <Route path="/reviews" element={
            <ProtectedRoute>
              <ReviewList />
            </ProtectedRoute>
          } />
          <Route path="/notifications" element={
            <ProtectedRoute>
              <NotificationList />
            </ProtectedRoute>
          } />

          {/* 管理员路由 */}
          <Route path="/admin/login" element={<AdminLogin />} />
          <Route path="/admin/dashboard" element={
            <AdminRoute>
              <AdminDashboard />
            </AdminRoute>
          } />
          <Route path="/admin/users" element={
            <AdminRoute>
              <AdminUserList />
            </AdminRoute>
          } />
          <Route path="/admin/services" element={
            <AdminRoute>
              <AdminServiceList />
            </AdminRoute>
          } />
          <Route path="/admin/orders" element={
            <AdminRoute>
              <AdminOrderList />
            </AdminRoute>
          } />

          {/* 404页面 */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Content>
    </Layout>
  );
}

function App() {
  return (
    <ConfigProvider locale={zhCN}>
      <AuthProvider>
        <AppContent />
      </AuthProvider>
    </ConfigProvider>
  );
}

export default App;