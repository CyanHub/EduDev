import React, { useState, useEffect } from 'react';
import { Layout, Menu, Card, Row, Col, Statistic, Table, Typography, Button, Spin, message } from 'antd';
import { 
  UserOutlined, 
  HomeOutlined, 
  ShoppingOutlined, 
  CommentOutlined, 
  BellOutlined,
  PieChartOutlined,
  TeamOutlined,
  ShopOutlined,
  OrderedListOutlined,
  LogoutOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { api } from '../../utils/api';

const { Header, Sider, Content } = Layout;
const { Title, Text } = Typography;

const AdminDashboard = () => {
  const { logout, currentUser } = useAuth();
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    totalUsers: 0,
    totalServices: 0,
    totalOrders: 0,
    totalRevenue: 0,
    pendingOrders: 0,
    completedOrders: 0
  });
  const [recentOrders, setRecentOrders] = useState([]);
  const [recentUsers, setRecentUsers] = useState([]);

  useEffect(() => {
    const fetchDashboardData = async () => {
      try {
        setLoading(true);
        // 获取仪表盘数据
        const response = await api.admin.getDashboard();
        if (response.data) {
          setStats({
            totalUsers: response.data.totalUsers || 0,
            totalServices: response.data.totalServices || 0,
            totalOrders: response.data.totalOrders || 0,
            totalRevenue: response.data.totalRevenue || 0,
            pendingOrders: response.data.pendingOrders || 0,
            completedOrders: response.data.completedOrders || 0
          });
          setRecentOrders(response.data.recentOrders || []);
          setRecentUsers(response.data.recentUsers || []);
        }
      } catch (error) {
        console.error('获取仪表盘数据失败:', error);
        message.error('获取仪表盘数据失败');
      } finally {
        setLoading(false);
      }
    };

    fetchDashboardData();
  }, []);

  const handleLogout = () => {
    logout();
    navigate('/admin/login');
    message.success('已成功退出登录');
  };

  const orderColumns = [
    {
      title: '订单ID',
      dataIndex: 'id',
      key: 'id',
      width: 80,
    },
    {
      title: '用户',
      dataIndex: 'userName',
      key: 'userName',
      width: 120,
    },
    {
      title: '服务',
      dataIndex: 'serviceName',
      key: 'serviceName',
      width: 150,
      ellipsis: true,
    },
    {
      title: '金额',
      dataIndex: 'amount',
      key: 'amount',
      width: 100,
      render: (amount) => `¥${amount.toFixed(2)}`,
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      width: 100,
      render: (status) => {
        let color = '';
        switch(status) {
          case 'PENDING':
            color = 'orange';
            status = '待处理';
            break;
          case 'CONFIRMED':
            color = 'blue';
            status = '已确认';
            break;
          case 'IN_PROGRESS':
            color = 'cyan';
            status = '进行中';
            break;
          case 'COMPLETED':
            color = 'green';
            status = '已完成';
            break;
          case 'CANCELLED':
            color = 'red';
            status = '已取消';
            break;
          default:
            color = 'default';
        }
        return <Text style={{ color }}>{status}</Text>;
      },
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      key: 'createdAt',
      width: 170,
    },
    {
      title: '操作',
      key: 'action',
      width: 100,
      render: (_, record) => (
        <Button type="link" size="small">
          <Link to={`/admin/orders/${record.id}`}>查看</Link>
        </Button>
      ),
    },
  ];

  const userColumns = [
    {
      title: '用户ID',
      dataIndex: 'id',
      key: 'id',
      width: 80,
    },
    {
      title: '用户名',
      dataIndex: 'username',
      key: 'username',
      width: 120,
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      key: 'email',
      width: 180,
      ellipsis: true,
    },
    {
      title: '电话',
      dataIndex: 'phone',
      key: 'phone',
      width: 130,
    },
    {
      title: '角色',
      dataIndex: 'role',
      key: 'role',
      width: 100,
      render: (role) => {
        return role === 'SERVICE_PROVIDER' ? '服务提供商' : '普通用户';
      },
    },
    {
      title: '注册时间',
      dataIndex: 'createdAt',
      key: 'createdAt',
      width: 170,
    },
    {
      title: '操作',
      key: 'action',
      width: 100,
      render: (_, record) => (
        <Button type="link" size="small">
          <Link to={`/admin/users/${record.id}`}>查看</Link>
        </Button>
      ),
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={setCollapsed}>
        <div className="logo" style={{ 
          height: '64px', 
          display: 'flex', 
          justifyContent: 'center', 
          alignItems: 'center',
          color: 'white',
          fontSize: collapsed ? '16px' : '20px',
          fontWeight: 'bold',
          margin: '16px 0'
        }}>
          {collapsed ? '宠管' : '宠物寄养管理'}
        </div>
        <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
          <Menu.Item key="1" icon={<PieChartOutlined />}>
            <Link to="/admin/dashboard">仪表盘</Link>
          </Menu.Item>
          <Menu.Item key="2" icon={<TeamOutlined />}>
            <Link to="/admin/users">用户管理</Link>
          </Menu.Item>
          <Menu.Item key="3" icon={<ShopOutlined />}>
            <Link to="/admin/services">服务管理</Link>
          </Menu.Item>
          <Menu.Item key="4" icon={<OrderedListOutlined />}>
            <Link to="/admin/orders">订单管理</Link>
          </Menu.Item>
          <Menu.Item key="5" icon={<LogoutOutlined />} onClick={handleLogout}>
            退出登录
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout className="site-layout">
        <Header className="site-layout-background" style={{ padding: 0, background: '#fff' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '0 24px' }}>
            <Title level={4} style={{ margin: 0 }}>管理员仪表盘</Title>
            <div>
              <Text strong style={{ marginRight: 8 }}>
                欢迎，{currentUser?.username || '管理员'}
              </Text>
              <Button type="link" onClick={handleLogout} icon={<LogoutOutlined />}>
                退出
              </Button>
            </div>
          </div>
        </Header>
        <Content style={{ margin: '16px' }}>
          {loading ? (
            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '80vh' }}>
              <Spin size="large" tip="加载中..." />
            </div>
          ) : (
            <>
              <Row gutter={[16, 16]}>
                <Col xs={24} sm={12} md={8} lg={6}>
                  <Card>
                    <Statistic 
                      title="总用户数" 
                      value={stats.totalUsers} 
                      prefix={<UserOutlined />} 
                    />
                  </Card>
                </Col>
                <Col xs={24} sm={12} md={8} lg={6}>
                  <Card>
                    <Statistic 
                      title="总服务数" 
                      value={stats.totalServices} 
                      prefix={<HomeOutlined />} 
                    />
                  </Card>
                </Col>
                <Col xs={24} sm={12} md={8} lg={6}>
                  <Card>
                    <Statistic 
                      title="总订单数" 
                      value={stats.totalOrders} 
                      prefix={<ShoppingOutlined />} 
                    />
                  </Card>
                </Col>
                <Col xs={24} sm={12} md={8} lg={6}>
                  <Card>
                    <Statistic 
                      title="总收入" 
                      value={stats.totalRevenue} 
                      precision={2}
                      prefix="¥" 
                    />
                  </Card>
                </Col>
                <Col xs={24} sm={12} md={12}>
                  <Card>
                    <Statistic 
                      title="待处理订单" 
                      value={stats.pendingOrders} 
                      valueStyle={{ color: '#faad14' }}
                    />
                  </Card>
                </Col>
                <Col xs={24} sm={12} md={12}>
                  <Card>
                    <Statistic 
                      title="已完成订单" 
                      value={stats.completedOrders} 
                      valueStyle={{ color: '#52c41a' }}
                    />
                  </Card>
                </Col>
              </Row>

              <Row gutter={[16, 16]} style={{ marginTop: 16 }}>
                <Col span={24}>
                  <Card title="最近订单" extra={<Link to="/admin/orders">查看全部</Link>}>
                    <Table 
                      columns={orderColumns} 
                      dataSource={recentOrders} 
                      rowKey="id" 
                      pagination={false}
                      scroll={{ x: 900 }}
                    />
                  </Card>
                </Col>
              </Row>

              <Row gutter={[16, 16]} style={{ marginTop: 16 }}>
                <Col span={24}>
                  <Card title="最近注册用户" extra={<Link to="/admin/users">查看全部</Link>}>
                    <Table 
                      columns={userColumns} 
                      dataSource={recentUsers} 
                      rowKey="id" 
                      pagination={false}
                      scroll={{ x: 900 }}
                    />
                  </Card>
                </Col>
              </Row>
            </>
          )}
        </Content>
      </Layout>
    </Layout>
  );
};

export default AdminDashboard;