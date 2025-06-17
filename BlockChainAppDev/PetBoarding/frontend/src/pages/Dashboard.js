import React, { useState, useEffect } from 'react';
import { Layout, Menu, Card, Statistic, Row, Col, Typography, Avatar, Button, List, Tag, Spin } from 'antd';
import { 
  UserOutlined, 
  PieChartOutlined, 
  ShoppingOutlined, 
  HeartOutlined,
  BellOutlined,
  HomeOutlined,
  PlusOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';

const { Header, Content, Sider } = Layout;
const { Title, Text } = Typography;

const Dashboard = () => {
  const { user } = useAuth();
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    totalPets: 0,
    activeOrders: 0,
    completedOrders: 0,
    reviews: 0
  });
  const [recentOrders, setRecentOrders] = useState([]);
  const [recentNotifications, setRecentNotifications] = useState([]);

  useEffect(() => {
    const fetchDashboardData = async () => {
      setLoading(true);
      try {
        // 获取用户统计数据
        const statsResponse = await api.user.getUserStats();
        if (statsResponse.data) {
          setStats(statsResponse.data);
        }

        // 获取最近订单
        const ordersResponse = await api.order.getRecentOrders();
        if (ordersResponse.data) {
          setRecentOrders(ordersResponse.data.slice(0, 5));
        }

        // 获取最近通知
        const notificationsResponse = await api.notification.getNotifications();
        if (notificationsResponse.data) {
          setRecentNotifications(notificationsResponse.data.slice(0, 5));
        }
      } catch (error) {
        console.error('获取仪表盘数据失败:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchDashboardData();
  }, []);

  const getOrderStatusTag = (status) => {
    const statusMap = {
      0: { color: 'blue', text: '待确认' },
      1: { color: 'processing', text: '进行中' },
      2: { color: 'success', text: '已完成' },
      3: { color: 'warning', text: '已取消' },
      4: { color: 'error', text: '已拒绝' }
    };
    const statusInfo = statusMap[status] || { color: 'default', text: '未知状态' };
    return <Tag color={statusInfo.color}>{statusInfo.text}</Tag>;
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['dashboard']}
          style={{ height: '100%', borderRight: 0 }}
        >
          <Menu.Item key="dashboard" icon={<PieChartOutlined />}>
            <Link to="/dashboard">仪表盘</Link>
          </Menu.Item>
          <Menu.Item key="pets" icon={<HomeOutlined />}>
            <Link to="/pets">我的宠物</Link>
          </Menu.Item>
          <Menu.Item key="services" icon={<ShoppingOutlined />}>
            <Link to="/services">寄养服务</Link>
          </Menu.Item>
          <Menu.Item key="orders" icon={<ShoppingOutlined />}>
            <Link to="/orders">我的订单</Link>
          </Menu.Item>
          <Menu.Item key="reviews" icon={<HeartOutlined />}>
            <Link to="/reviews">我的评价</Link>
          </Menu.Item>
          <Menu.Item key="notifications" icon={<BellOutlined />}>
            <Link to="/notifications">通知消息</Link>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout style={{ padding: '0 24px 24px' }}>
        <Header style={{ background: '#fff', padding: 0, margin: '16px 0' }}>
          <div style={{ display: 'flex', alignItems: 'center', padding: '0 24px' }}>
            <Avatar size="large" icon={<UserOutlined />} />
            <div style={{ marginLeft: 16 }}>
              <Title level={4} style={{ margin: 0 }}>{user?.username || '用户'}</Title>
              <Text type="secondary">{user?.email || ''}</Text>
            </div>
          </div>
        </Header>
        <Content
          className="site-layout-background"
          style={{ padding: 24, margin: 0, minHeight: 280 }}
        >
          {loading ? (
            <div style={{ textAlign: 'center', padding: '50px 0' }}>
              <Spin size="large" />
            </div>
          ) : (
            <>
              <Row gutter={16} style={{ marginBottom: 24 }}>
                <Col span={6}>
                  <Card>
                    <Statistic 
                      title="我的宠物" 
                      value={stats.totalPets} 
                      prefix={<HomeOutlined />} 
                    />
                    <Button 
                      type="link" 
                      style={{ padding: 0, marginTop: 8 }}
                      onClick={() => navigate('/pets')}
                    >
                      查看详情
                    </Button>
                  </Card>
                </Col>
                <Col span={6}>
                  <Card>
                    <Statistic 
                      title="进行中订单" 
                      value={stats.activeOrders} 
                      prefix={<ShoppingOutlined />} 
                    />
                    <Button 
                      type="link" 
                      style={{ padding: 0, marginTop: 8 }}
                      onClick={() => navigate('/orders?status=1')}
                    >
                      查看详情
                    </Button>
                  </Card>
                </Col>
                <Col span={6}>
                  <Card>
                    <Statistic 
                      title="已完成订单" 
                      value={stats.completedOrders} 
                      prefix={<ShoppingOutlined />} 
                    />
                    <Button 
                      type="link" 
                      style={{ padding: 0, marginTop: 8 }}
                      onClick={() => navigate('/orders?status=2')}
                    >
                      查看详情
                    </Button>
                  </Card>
                </Col>
                <Col span={6}>
                  <Card>
                    <Statistic 
                      title="我的评价" 
                      value={stats.reviews} 
                      prefix={<HeartOutlined />} 
                    />
                    <Button 
                      type="link" 
                      style={{ padding: 0, marginTop: 8 }}
                      onClick={() => navigate('/reviews')}
                    >
                      查看详情
                    </Button>
                  </Card>
                </Col>
              </Row>

              <Row gutter={16}>
                <Col span={12}>
                  <Card 
                    title="最近订单" 
                    extra={<Link to="/orders">查看全部</Link>}
                  >
                    <List
                      itemLayout="horizontal"
                      dataSource={recentOrders}
                      locale={{ emptyText: '暂无订单' }}
                      renderItem={item => (
                        <List.Item
                          actions={[
                            <Link to={`/orders/${item.order_id}`}>查看详情</Link>
                          ]}
                        >
                          <List.Item.Meta
                            title={
                              <div>
                                <span style={{ marginRight: 8 }}>{item.service_name}</span>
                                {getOrderStatusTag(item.status)}
                              </div>
                            }
                            description={`订单号: ${item.order_id} | 宠物: ${item.pet_name} | 创建时间: ${new Date(item.created_at).toLocaleString()}`}
                          />
                          <div>¥{item.total_amount}</div>
                        </List.Item>
                      )}
                    />
                  </Card>
                </Col>
                <Col span={12}>
                  <Card 
                    title="最近通知" 
                    extra={<Link to="/notifications">查看全部</Link>}
                  >
                    <List
                      itemLayout="horizontal"
                      dataSource={recentNotifications}
                      locale={{ emptyText: '暂无通知' }}
                      renderItem={item => (
                        <List.Item>
                          <List.Item.Meta
                            title={item.title}
                            description={
                              <>
                                <div>{item.content}</div>
                                <div style={{ fontSize: 12, color: '#999' }}>
                                  {new Date(item.created_at).toLocaleString()}
                                </div>
                              </>
                            }
                          />
                        </List.Item>
                      )}
                    />
                  </Card>
                </Col>
              </Row>

              <Row style={{ marginTop: 16 }}>
                <Col span={24}>
                  <Card>
                    <div style={{ textAlign: 'center' }}>
                      <Button 
                        type="primary" 
                        icon={<PlusOutlined />}
                        onClick={() => navigate('/services')}
                      >
                        预约新的寄养服务
                      </Button>
                    </div>
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

export default Dashboard;