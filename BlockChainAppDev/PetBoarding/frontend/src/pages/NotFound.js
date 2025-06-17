import React from 'react';
import { Result, Button } from 'antd';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const NotFound = () => {
  const { currentUser } = useAuth();
  
  // 根据用户角色决定返回的首页路径
  const getHomePath = () => {
    if (!currentUser) return '/login';
    if (currentUser.role === 'ADMIN') return '/admin/dashboard';
    return '/dashboard';
  };

  return (
    <div style={{ 
      display: 'flex', 
      justifyContent: 'center', 
      alignItems: 'center', 
      minHeight: '100vh',
      background: '#f0f2f5'
    }}>
      <Result
        status="404"
        title="404"
        subTitle="抱歉，您访问的页面不存在。"
        extra={
          <Button type="primary">
            <Link to={getHomePath()}>返回首页</Link>
          </Button>
        }
      />
    </div>
  );
};

export default NotFound;