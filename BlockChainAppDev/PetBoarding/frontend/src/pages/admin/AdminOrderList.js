import React, { useState, useEffect } from 'react';
import { Layout, Menu, Table, Card, Button, Input, Select, Modal, Form, message, Tag, Space, Spin, Popconfirm, Steps, Descriptions, Row, Col } from 'antd';
import { 
  PieChartOutlined, 
  TeamOutlined, 
  ShopOutlined, 
  OrderedListOutlined, 
  LogoutOutlined,
  SearchOutlined,
  EyeOutlined,
  CheckOutlined,
  CloseOutlined,
  ExclamationCircleOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { api } from '../../utils/api';

const { Header, Sider, Content } = Layout;
const { Option } = Select;
const { Step } = Steps;

const AdminOrderList = () => {
  const { logout, currentUser } = useAuth();
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const [loading, setLoading] = useState(true);
  const [orders, setOrders] = useState([]);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });
  const [searchForm] = Form.useForm();
  const [detailModalVisible, setDetailModalVisible] = useState(false);
  const [currentOrder, setCurrentOrder] = useState(null);
  const [searchParams, setSearchParams] = useState({
    orderId: '',
    userId: '',
    serviceId: '',
    status: ''
  });

  const fetchOrders = async (page = 1, pageSize = 10, params = {}) => {
    try {
      setLoading(true);
      const response = await api.admin.getOrders({
        page,
        size: pageSize,
        ...params
      });
      
      if (response.data) {
        setOrders(response.data.content || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.totalElements || 0
        });
      }
    } catch (error) {
      console.error('获取订单列表失败:', error);
      message.error('获取订单列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchOrders(pagination.current, pagination.pageSize, searchParams);
  }, []);

  const handleTableChange = (pagination) => {
    fetchOrders(pagination.current, pagination.pageSize, searchParams);
  };

  const handleSearch = (values) => {
    const params = {
      orderId: values.orderId || '',
      userId: values.userId || '',
      serviceId: values.serviceId || '',
      status: values.status || ''
    };
    setSearchParams(params);
    fetchOrders(1, pagination.pageSize, params);
  };

  const handleReset = () => {
    searchForm.resetFields();
    const params = {
      orderId: '',
      userId: '',
      serviceId: '',
      status: ''
    };
    setSearchParams(params);
    fetchOrders(1, pagination.pageSize, params);
  };

  const showDetailModal = (order) => {
    setCurrentOrder(order);
    setDetailModalVisible(true);
  };

  const handleUpdateOrderStatus = async (orderId, newStatus) => {
    try {
      await api.admin.updateOrderStatus(orderId, { status: newStatus });
      message.success('订单状态更新成功');
      fetchOrders(pagination.current, pagination.pageSize, searchParams);
      if (detailModalVisible) {
        // 如果详情模态框打开，更新当前订单信息
        const updatedOrder = { ...currentOrder, status: newStatus };
        setCurrentOrder(updatedOrder);
      }
    } catch (error) {
      console.error('更新订单状态失败:', error);
      message.error('更新订单状态失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/admin/login');
    message.success('已成功退出登录');
  };

  // 获取订单状态对应的步骤
  const getOrderStep = (status) => {
    switch (status) {
      case 'PENDING':
        return 0;
      case 'CONFIRMED':
        return 1;
      case 'IN_PROGRESS':
        return 2;
      case 'COMPLETED':
        return 3;
      case 'CANCELLED':
        return -1; // 取消状态
      default:
        return 0;
    }
  };

  // 获取订单状态中文名称和颜色
  const getOrderStatusInfo = (status) => {
    switch (status) {
      case 'PENDING':
        return { text: '待确认', color: 'orange' };
      case 'CONFIRMED':
        return { text: '已确认', color: 'blue' };
      case 'IN_PROGRESS':
        return { text: '进行中', color: 'cyan' };
      case 'COMPLETED':
        return { text: '已完成', color: 'green' };
      case 'CANCELLED':
        return { text: '已取消', color: 'red' };
      default:
        return { text: '未知状态', color: 'default' };
    }
  };

  const columns = [
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
      title: '宠物',
      dataIndex: 'petName',
      key: 'petName',
      width: 120,
    },
    {
      title: '开始日期',
      dataIndex: 'startDate',
      key: 'startDate',
      width: 120,
    },
    {
      title: '结束日期',
      dataIndex: 'endDate',
      key: 'endDate',
      width: 120,
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
        const statusInfo = getOrderStatusInfo(status);
        return <Tag color={statusInfo.color}>{statusInfo.text}</Tag>;
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
      width: 200,
      render: (_, record) => (
        <Space size="small">
          <Button 
            type="primary" 
            icon={<EyeOutlined />} 
            size="small" 
            onClick={() => showDetailModal(record)}
          >
            详情
          </Button>
          {record.status === 'PENDING' && (
            <Button 
              type="primary" 
              icon={<CheckOutlined />} 
              size="small" 
              onClick={() => handleUpdateOrderStatus(record.id, 'CONFIRMED')}
            >
              确认
            </Button>
          )}
          {record.status === 'CONFIRMED' && (
            <Button 
              type="primary" 
              icon={<CheckOutlined />} 
              size="small" 
              onClick={() => handleUpdateOrderStatus(record.id, 'IN_PROGRESS')}
            >
              开始服务
            </Button>
          )}
          {record.status === 'IN_PROGRESS' && (
            <Button 
              type="primary" 
              icon={<CheckOutlined />} 
              size="small" 
              onClick={() => handleUpdateOrderStatus(record.id, 'COMPLETED')}
            >
              完成
            </Button>
          )}
          {(record.status === 'PENDING' || record.status === 'CONFIRMED') && (
            <Popconfirm
              title="确定要取消此订单吗？"
              icon={<ExclamationCircleOutlined style={{ color: 'red' }} />}
              onConfirm={() => handleUpdateOrderStatus(record.id, 'CANCELLED')}
              okText="确定"
              cancelText="取消"
            >
              <Button 
                danger 
                icon={<CloseOutlined />} 
                size="small"
              >
                取消
              </Button>
            </Popconfirm>
          )}
        </Space>
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
        <Menu theme="dark" defaultSelectedKeys={['4']} mode="inline">
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
            <h2 style={{ margin: 0 }}>订单管理</h2>
            <div>
              <span style={{ marginRight: 12 }}>欢迎，{currentUser?.username || '管理员'}</span>
              <Button type="link" onClick={handleLogout} icon={<LogoutOutlined />}>
                退出
              </Button>
            </div>
          </div>
        </Header>
        <Content style={{ margin: '16px' }}>
          <Card>
            <Form
              form={searchForm}
              layout="inline"
              onFinish={handleSearch}
              style={{ marginBottom: 16 }}
            >
              <Form.Item name="orderId">
                <Input 
                  placeholder="订单ID" 
                  allowClear 
                />
              </Form.Item>
              <Form.Item name="userId">
                <Input 
                  placeholder="用户ID" 
                  allowClear 
                />
              </Form.Item>
              <Form.Item name="serviceId">
                <Input 
                  placeholder="服务ID" 
                  allowClear 
                />
              </Form.Item>
              <Form.Item name="status">
                <Select 
                  placeholder="订单状态" 
                  style={{ width: 120 }} 
                  allowClear
                >
                  <Option value="PENDING">待确认</Option>
                  <Option value="CONFIRMED">已确认</Option>
                  <Option value="IN_PROGRESS">进行中</Option>
                  <Option value="COMPLETED">已完成</Option>
                  <Option value="CANCELLED">已取消</Option>
                </Select>
              </Form.Item>
              <Form.Item>
                <Button 
                  type="primary" 
                  htmlType="submit" 
                  icon={<SearchOutlined />}
                >
                  搜索
                </Button>
              </Form.Item>
              <Form.Item>
                <Button onClick={handleReset}>重置</Button>
              </Form.Item>
            </Form>

            <Table 
              columns={columns} 
              dataSource={orders} 
              rowKey="id" 
              pagination={pagination}
              loading={loading}
              onChange={handleTableChange}
              scroll={{ x: 1300 }}
            />
          </Card>

          {/* 订单详情模态框 */}
          <Modal
            title="订单详情"
            visible={detailModalVisible}
            onCancel={() => setDetailModalVisible(false)}
            footer={[
              <Button key="close" onClick={() => setDetailModalVisible(false)}>
                关闭
              </Button>,
              currentOrder?.status === 'PENDING' && (
                <Button 
                  key="confirm" 
                  type="primary" 
                  onClick={() => handleUpdateOrderStatus(currentOrder.id, 'CONFIRMED')}
                >
                  确认订单
                </Button>
              ),
              currentOrder?.status === 'CONFIRMED' && (
                <Button 
                  key="start" 
                  type="primary" 
                  onClick={() => handleUpdateOrderStatus(currentOrder.id, 'IN_PROGRESS')}
                >
                  开始服务
                </Button>
              ),
              currentOrder?.status === 'IN_PROGRESS' && (
                <Button 
                  key="complete" 
                  type="primary" 
                  onClick={() => handleUpdateOrderStatus(currentOrder.id, 'COMPLETED')}
                >
                  完成服务
                </Button>
              ),
              (currentOrder?.status === 'PENDING' || currentOrder?.status === 'CONFIRMED') && (
                <Button 
                  key="cancel" 
                  danger 
                  onClick={() => {
                    Modal.confirm({
                      title: '确认取消订单',
                      icon: <ExclamationCircleOutlined />,
                      content: '确定要取消此订单吗？',
                      okText: '确定',
                      cancelText: '取消',
                      onOk: () => handleUpdateOrderStatus(currentOrder.id, 'CANCELLED')
                    });
                  }}
                >
                  取消订单
                </Button>
              )
            ]}
            width={800}
          >
            {currentOrder && (
              <div>
                <Row gutter={[16, 16]}>
                  <Col span={24}>
                    <Steps current={getOrderStep(currentOrder.status)} status={currentOrder.status === 'CANCELLED' ? 'error' : 'process'}>
                      <Step title="待确认" description="订单已提交" />
                      <Step title="已确认" description="订单已确认" />
                      <Step title="进行中" description="服务进行中" />
                      <Step title="已完成" description="服务已完成" />
                    </Steps>
                  </Col>
                </Row>

                <Descriptions title="订单信息" bordered style={{ marginTop: 24 }}>
                  <Descriptions.Item label="订单ID" span={3}>{currentOrder.id}</Descriptions.Item>
                  <Descriptions.Item label="订单状态" span={3}>
                    <Tag color={getOrderStatusInfo(currentOrder.status).color}>
                      {getOrderStatusInfo(currentOrder.status).text}
                    </Tag>
                  </Descriptions.Item>
                  <Descriptions.Item label="创建时间" span={3}>{currentOrder.createdAt}</Descriptions.Item>
                  <Descriptions.Item label="更新时间" span={3}>{currentOrder.updatedAt}</Descriptions.Item>
                </Descriptions>

                <Descriptions title="服务信息" bordered style={{ marginTop: 24 }}>
                  <Descriptions.Item label="服务ID">{currentOrder.serviceId}</Descriptions.Item>
                  <Descriptions.Item label="服务名称" span={2}>{currentOrder.serviceName}</Descriptions.Item>
                  <Descriptions.Item label="服务提供商">{currentOrder.providerName}</Descriptions.Item>
                  <Descriptions.Item label="开始日期">{currentOrder.startDate}</Descriptions.Item>
                  <Descriptions.Item label="结束日期">{currentOrder.endDate}</Descriptions.Item>
                  <Descriptions.Item label="服务天数">{currentOrder.days}</Descriptions.Item>
                  <Descriptions.Item label="单价">{`¥${currentOrder.price?.toFixed(2)}/天`}</Descriptions.Item>
                  <Descriptions.Item label="总金额" span={1}>{`¥${currentOrder.amount?.toFixed(2)}`}</Descriptions.Item>
                </Descriptions>

                <Descriptions title="用户信息" bordered style={{ marginTop: 24 }}>
                  <Descriptions.Item label="用户ID">{currentOrder.userId}</Descriptions.Item>
                  <Descriptions.Item label="用户名">{currentOrder.userName}</Descriptions.Item>
                  <Descriptions.Item label="联系电话">{currentOrder.userPhone}</Descriptions.Item>
                </Descriptions>

                <Descriptions title="宠物信息" bordered style={{ marginTop: 24 }}>
                  <Descriptions.Item label="宠物ID">{currentOrder.petId}</Descriptions.Item>
                  <Descriptions.Item label="宠物名称">{currentOrder.petName}</Descriptions.Item>
                  <Descriptions.Item label="宠物类型">{currentOrder.petType}</Descriptions.Item>
                  <Descriptions.Item label="宠物品种">{currentOrder.petBreed}</Descriptions.Item>
                  <Descriptions.Item label="宠物年龄">{currentOrder.petAge}</Descriptions.Item>
                  <Descriptions.Item label="宠物性别">{currentOrder.petGender === 'MALE' ? '公' : '母'}</Descriptions.Item>
                </Descriptions>

                {currentOrder.notes && (
                  <Descriptions title="备注信息" bordered style={{ marginTop: 24 }}>
                    <Descriptions.Item span={3}>{currentOrder.notes}</Descriptions.Item>
                  </Descriptions>
                )}
              </div>
            )}
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default AdminOrderList;