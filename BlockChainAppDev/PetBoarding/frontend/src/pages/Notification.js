import React, { useState, useEffect } from 'react';
import { 
  Layout, Menu, Card, List, Badge, Button, 
  Tabs, Empty, Spin, message, Typography, Tag, Modal 
} from 'antd';
import { 
  UserOutlined, PieChartOutlined, ShoppingOutlined, 
  HeartOutlined, BellOutlined, HomeOutlined, 
  CheckCircleOutlined, ExclamationCircleOutlined, 
  InfoCircleOutlined, DeleteOutlined, EyeOutlined 
} from '@ant-design/icons';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { TabPane } = Tabs;
const { Title, Text, Paragraph } = Typography;
const { confirm } = Modal;

const Notification = () => {
  const { user } = useAuth();
  const [notifications, setNotifications] = useState([]);
  const [loading, setLoading] = useState(true);
  const [unreadCount, setUnreadCount] = useState(0);
  const [detailVisible, setDetailVisible] = useState(false);
  const [currentNotification, setCurrentNotification] = useState(null);

  // 获取通知列表
  const fetchNotifications = async () => {
    setLoading(true);
    try {
      const response = await api.notification.getNotifications();
      if (response.data) {
        setNotifications(response.data);
        // 计算未读通知数量
        const unread = response.data.filter(item => !item.is_read).length;
        setUnreadCount(unread);
      }
    } catch (error) {
      console.error('获取通知列表失败:', error);
      message.error('获取通知列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchNotifications();
  }, []);

  // 标记通知为已读
  const markAsRead = async (notificationId) => {
    try {
      await api.notification.markAsRead(notificationId);
      // 更新本地通知状态
      setNotifications(prevNotifications => 
        prevNotifications.map(item => 
          item.notification_id === notificationId 
            ? { ...item, is_read: true } 
            : item
        )
      );
      // 更新未读计数
      setUnreadCount(prevCount => Math.max(0, prevCount - 1));
    } catch (error) {
      console.error('标记通知已读失败:', error);
    }
  };

  // 标记所有通知为已读
  const markAllAsRead = async () => {
    try {
      await api.notification.markAllAsRead();
      // 更新本地通知状态
      setNotifications(prevNotifications => 
        prevNotifications.map(item => ({ ...item, is_read: true }))
      );
      // 更新未读计数
      setUnreadCount(0);
      message.success('已将所有通知标记为已读');
    } catch (error) {
      console.error('标记所有通知已读失败:', error);
      message.error('操作失败');
    }
  };

  // 删除通知
  const deleteNotification = async (notificationId) => {
    confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: '确定要删除这条通知吗？',
      onOk: async () => {
        try {
          await api.notification.deleteNotification(notificationId);
          // 从列表中移除
          setNotifications(prevNotifications => 
            prevNotifications.filter(item => item.notification_id !== notificationId)
          );
          // 如果是未读通知，更新未读计数
          const notification = notifications.find(item => item.notification_id === notificationId);
          if (notification && !notification.is_read) {
            setUnreadCount(prevCount => Math.max(0, prevCount - 1));
          }
          message.success('通知删除成功');
        } catch (error) {
          console.error('删除通知失败:', error);
          message.error('删除通知失败');
        }
      },
    });
  };

  // 查看通知详情
  const viewNotificationDetail = (notification) => {
    setCurrentNotification(notification);
    setDetailVisible(true);
    
    // 如果是未读通知，标记为已读
    if (!notification.is_read) {
      markAsRead(notification.notification_id);
    }
  };

  // 关闭详情模态框
  const handleDetailClose = () => {
    setDetailVisible(false);
  };

  // 获取通知图标
  const getNotificationIcon = (type) => {
    switch (type) {
      case 'system':
        return <InfoCircleOutlined style={{ color: '#1890ff' }} />;
      case 'order':
        return <ShoppingOutlined style={{ color: '#52c41a' }} />;
      case 'service':
        return <HomeOutlined style={{ color: '#faad14' }} />;
      case 'review':
        return <HeartOutlined style={{ color: '#f5222d' }} />;
      default:
        return <BellOutlined />;
    }
  };

  // 获取通知类型标签
  const getNotificationTypeTag = (type) => {
    const typeMap = {
      'system': { color: 'blue', text: '系统通知' },
      'order': { color: 'green', text: '订单通知' },
      'service': { color: 'orange', text: '服务通知' },
      'review': { color: 'red', text: '评价通知' },
    };
    const typeInfo = typeMap[type] || { color: 'default', text: '其他通知' };
    return <Tag color={typeInfo.color}>{typeInfo.text}</Tag>;
  };

  // 渲染通知列表项
  const renderNotificationItem = (item) => (
    <List.Item
      key={item.notification_id}
      actions={[
        <Button 
          type="link" 
          icon={<EyeOutlined />} 
          onClick={() => viewNotificationDetail(item)}
        >
          查看
        </Button>,
        <Button 
          type="link" 
          danger 
          icon={<DeleteOutlined />} 
          onClick={() => deleteNotification(item.notification_id)}
        >
          删除
        </Button>,
      ]}
    >
      <List.Item.Meta
        avatar={
          <Badge dot={!item.is_read}>
            {getNotificationIcon(item.type)}
          </Badge>
        }
        title={
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span style={{ fontWeight: !item.is_read ? 'bold' : 'normal' }}>
              {item.title}
            </span>
            <div>
              {getNotificationTypeTag(item.type)}
            </div>
          </div>
        }
        description={
          <>
            <div style={{ marginBottom: 8 }}>
              <Paragraph ellipsis={{ rows: 2 }}>
                {item.content}
              </Paragraph>
            </div>
            <div style={{ fontSize: 12, color: '#999' }}>
              {moment(item.created_at).format('YYYY-MM-DD HH:mm:ss')}
            </div>
          </>
        }
      />
    </List.Item>
  );

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['notifications']}
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
        <Content
          className="site-layout-background"
          style={{ padding: 24, margin: '16px 0', minHeight: 280 }}
        >
          <Card 
            title={
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <span>通知消息</span>
                {unreadCount > 0 && (
                  <Button type="link" onClick={markAllAsRead}>
                    全部标为已读
                  </Button>
                )}
              </div>
            }
          >
            <Tabs defaultActiveKey="all">
              <TabPane 
                tab={
                  <span>
                    全部通知
                    {unreadCount > 0 && (
                      <Badge 
                        count={unreadCount} 
                        style={{ marginLeft: 8 }} 
                        size="small" 
                      />
                    )}
                  </span>
                } 
                key="all"
              >
                {loading ? (
                  <div style={{ textAlign: 'center', padding: '50px 0' }}>
                    <Spin size="large" />
                  </div>
                ) : notifications.length > 0 ? (
                  <List
                    itemLayout="vertical"
                    dataSource={notifications}
                    pagination={{
                      pageSize: 10,
                      showTotal: (total) => `共 ${total} 条通知`,
                    }}
                    renderItem={renderNotificationItem}
                  />
                ) : (
                  <Empty description="暂无通知消息" />
                )}
              </TabPane>
              
              <TabPane tab="未读通知" key="unread">
                {loading ? (
                  <div style={{ textAlign: 'center', padding: '50px 0' }}>
                    <Spin size="large" />
                  </div>
                ) : notifications.filter(item => !item.is_read).length > 0 ? (
                  <List
                    itemLayout="vertical"
                    dataSource={notifications.filter(item => !item.is_read)}
                    pagination={{
                      pageSize: 10,
                      showTotal: (total) => `共 ${total} 条未读通知`,
                    }}
                    renderItem={renderNotificationItem}
                  />
                ) : (
                  <Empty description="暂无未读通知" />
                )}
              </TabPane>
              
              <TabPane tab="系统通知" key="system">
                {loading ? (
                  <div style={{ textAlign: 'center', padding: '50px 0' }}>
                    <Spin size="large" />
                  </div>
                ) : notifications.filter(item => item.type === 'system').length > 0 ? (
                  <List
                    itemLayout="vertical"
                    dataSource={notifications.filter(item => item.type === 'system')}
                    pagination={{
                      pageSize: 10,
                      showTotal: (total) => `共 ${total} 条系统通知`,
                    }}
                    renderItem={renderNotificationItem}
                  />
                ) : (
                  <Empty description="暂无系统通知" />
                )}
              </TabPane>
              
              <TabPane tab="订单通知" key="order">
                {loading ? (
                  <div style={{ textAlign: 'center', padding: '50px 0' }}>
                    <Spin size="large" />
                  </div>
                ) : notifications.filter(item => item.type === 'order').length > 0 ? (
                  <List
                    itemLayout="vertical"
                    dataSource={notifications.filter(item => item.type === 'order')}
                    pagination={{
                      pageSize: 10,
                      showTotal: (total) => `共 ${total} 条订单通知`,
                    }}
                    renderItem={renderNotificationItem}
                  />
                ) : (
                  <Empty description="暂无订单通知" />
                )}
              </TabPane>
            </Tabs>
          </Card>

          {/* 通知详情模态框 */}
          <Modal
            title="通知详情"
            visible={detailVisible}
            onCancel={handleDetailClose}
            footer={[
              <Button key="back" onClick={handleDetailClose}>
                关闭
              </Button>,
            ]}
          >
            {currentNotification && (
              <div>
                <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                  <Title level={4}>{currentNotification.title}</Title>
                  {getNotificationTypeTag(currentNotification.type)}
                </div>
                
                <div style={{ marginBottom: 16 }}>
                  <Text type="secondary">
                    {moment(currentNotification.created_at).format('YYYY-MM-DD HH:mm:ss')}
                  </Text>
                </div>
                
                <div style={{ marginBottom: 16 }}>
                  <Paragraph style={{ fontSize: 16 }}>
                    {currentNotification.content}
                  </Paragraph>
                </div>
                
                {currentNotification.link && (
                  <div>
                    <Button type="primary">
                      <Link to={currentNotification.link}>查看详情</Link>
                    </Button>
                  </div>
                )}
              </div>
            )}
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default Notification;