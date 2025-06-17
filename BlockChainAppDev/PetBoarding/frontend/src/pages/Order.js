import React, { useState, useEffect } from 'react';
import { 
  Layout, Menu, Card, Button, Table, Tag, Modal, Form, 
  Input, Select, DatePicker, Descriptions, Divider, 
  Steps, message, Spin, Typography, Row, Col, Statistic 
} from 'antd';
import { 
  UserOutlined, PieChartOutlined, ShoppingOutlined, 
  HeartOutlined, BellOutlined, HomeOutlined, 
  CheckCircleOutlined, CloseCircleOutlined, 
  ExclamationCircleOutlined, ClockCircleOutlined,
  DollarOutlined, MessageOutlined
} from '@ant-design/icons';
import { Link, useNavigate, useParams, useLocation } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { Option } = Select;
const { Step } = Steps;
const { Title, Text, Paragraph } = Typography;
const { confirm } = Modal;

const Order = () => {
  const { user } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const { orderId } = useParams();
  const [form] = Form.useForm();
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [currentOrder, setCurrentOrder] = useState(null);
  const [detailVisible, setDetailVisible] = useState(false);
  const [reviewVisible, setReviewVisible] = useState(false);
  const [cancelVisible, setCancelVisible] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [statusFilter, setStatusFilter] = useState(null);

  // 从URL获取状态过滤器
  useEffect(() => {
    const params = new URLSearchParams(location.search);
    const status = params.get('status');
    if (status) {
      setStatusFilter(parseInt(status));
    }
  }, [location]);

  // 获取订单列表
  const fetchOrders = async () => {
    setLoading(true);
    try {
      const params = {};
      if (statusFilter !== null) {
        params.status = statusFilter;
      }
      
      const response = await api.order.getOrders(params);
      if (response.data) {
        setOrders(response.data);
      }
    } catch (error) {
      console.error('获取订单列表失败:', error);
      message.error('获取订单列表失败');
    } finally {
      setLoading(false);
    }
  };

  // 获取订单详情
  const fetchOrderDetail = async (id) => {
    setLoading(true);
    try {
      const response = await api.order.getOrderDetail(id);
      if (response.data) {
        setCurrentOrder(response.data);
        setDetailVisible(true);
      }
    } catch (error) {
      console.error('获取订单详情失败:', error);
      message.error('获取订单详情失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (orderId) {
      fetchOrderDetail(orderId);
    } else {
      fetchOrders();
    }
  }, [orderId, statusFilter]);

  // 关闭详情模态框
  const handleDetailClose = () => {
    setDetailVisible(false);
    if (orderId) {
      navigate('/orders');
    }
  };

  // 打开评价模态框
  const showReviewModal = (order) => {
    setCurrentOrder(order);
    form.resetFields();
    setReviewVisible(true);
  };

  // 关闭评价模态框
  const handleReviewClose = () => {
    setReviewVisible(false);
  };

  // 打开取消订单模态框
  const showCancelModal = (order) => {
    setCurrentOrder(order);
    form.resetFields();
    setCancelVisible(true);
  };

  // 关闭取消订单模态框
  const handleCancelClose = () => {
    setCancelVisible(false);
  };

  // 提交评价
  const handleSubmitReview = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      
      const reviewData = {
        order_id: currentOrder.order_id,
        service_id: currentOrder.service_id,
        rating: values.rating,
        content: values.content,
      };
      
      await api.review.addReview(reviewData);
      message.success('评价提交成功');
      setReviewVisible(false);
      fetchOrders();
    } catch (error) {
      console.error('提交评价失败:', error);
      message.error('提交评价失败');
    } finally {
      setSubmitting(false);
    }
  };

  // 取消订单
  const handleCancelOrder = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      
      await api.order.cancelOrder(currentOrder.order_id, {
        cancel_reason: values.cancel_reason
      });
      
      message.success('订单取消成功');
      setCancelVisible(false);
      fetchOrders();
    } catch (error) {
      console.error('取消订单失败:', error);
      message.error('取消订单失败');
    } finally {
      setSubmitting(false);
    }
  };

  // 支付订单
  const handlePayOrder = (order) => {
    confirm({
      title: '确认支付',
      icon: <ExclamationCircleOutlined />,
      content: `确定要支付订单 ${order.order_id} 吗？金额：¥${order.total_amount}`,
      onOk: async () => {
        try {
          await api.order.payOrder(order.order_id);
          message.success('支付成功');
          fetchOrders();
        } catch (error) {
          console.error('支付失败:', error);
          message.error('支付失败');
        }
      },
    });
  };

  // 获取订单状态标签
  const getStatusTag = (status) => {
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

  // 获取订单状态步骤
  const getOrderSteps = (order) => {
    const steps = [
      { title: '创建订单', icon: <ClockCircleOutlined /> },
      { title: '服务确认', icon: <CheckCircleOutlined /> },
      { title: '服务中', icon: <ClockCircleOutlined /> },
      { title: '服务完成', icon: <CheckCircleOutlined /> },
    ];
    
    let current = 0;
    
    switch (order.status) {
      case 0: // 待确认
        current = 0;
        break;
      case 1: // 进行中
        current = 2;
        break;
      case 2: // 已完成
        current = 3;
        break;
      case 3: // 已取消
      case 4: // 已拒绝
        current = 0;
        break;
      default:
        current = 0;
    }
    
    return (
      <Steps current={current} size="small">
        {steps.map((step, index) => (
          <Step key={index} title={step.title} icon={step.icon} />
        ))}
      </Steps>
    );
  };

  // 表格列定义
  const columns = [
    {
      title: '订单号',
      dataIndex: 'order_id',
      key: 'order_id',
      render: (text) => <a onClick={() => fetchOrderDetail(text)}>{text}</a>,
    },
    {
      title: '服务名称',
      dataIndex: 'service_name',
      key: 'service_name',
    },
    {
      title: '宠物名称',
      dataIndex: 'pet_name',
      key: 'pet_name',
    },
    {
      title: '开始日期',
      dataIndex: 'start_date',
      key: 'start_date',
      render: (text) => text ? moment(text).format('YYYY-MM-DD') : '-',
    },
    {
      title: '结束日期',
      dataIndex: 'end_date',
      key: 'end_date',
      render: (text) => text ? moment(text).format('YYYY-MM-DD') : '-',
    },
    {
      title: '总金额',
      dataIndex: 'total_amount',
      key: 'total_amount',
      render: (text) => `¥${text}`,
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (status) => getStatusTag(status),
      filters: [
        { text: '待确认', value: 0 },
        { text: '进行中', value: 1 },
        { text: '已完成', value: 2 },
        { text: '已取消', value: 3 },
        { text: '已拒绝', value: 4 },
      ],
      filteredValue: statusFilter !== null ? [statusFilter] : null,
      onFilter: (value, record) => record.status === value,
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (text) => text ? moment(text).format('YYYY-MM-DD HH:mm') : '-',
    },
    {
      title: '操作',
      key: 'action',
      render: (_, record) => {
        return (
          <>
            <Button 
              type="link" 
              onClick={() => fetchOrderDetail(record.order_id)}
            >
              详情
            </Button>
            
            {record.status === 0 && (
              <Button 
                type="link" 
                danger 
                onClick={() => showCancelModal(record)}
              >
                取消
              </Button>
            )}
            
            {record.status === 2 && !record.is_reviewed && (
              <Button 
                type="link" 
                onClick={() => showReviewModal(record)}
              >
                评价
              </Button>
            )}
          </>
        );
      },
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['orders']}
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
          <Card title="我的订单">
            {loading && !detailVisible ? (
              <div style={{ textAlign: 'center', padding: '50px 0' }}>
                <Spin size="large" />
              </div>
            ) : (
              <Table 
                columns={columns} 
                dataSource={orders} 
                rowKey="order_id" 
                pagination={{ pageSize: 10 }}
                locale={{ emptyText: '暂无订单' }}
                onChange={(pagination, filters) => {
                  if (filters.status && filters.status.length > 0) {
                    setStatusFilter(filters.status[0]);
                  } else {
                    setStatusFilter(null);
                  }
                }}
              />
            )}
          </Card>

          {/* 订单详情模态框 */}
          <Modal
            title="订单详情"
            visible={detailVisible}
            onCancel={handleDetailClose}
            footer={[
              <Button key="back" onClick={handleDetailClose}>
                关闭
              </Button>,
              currentOrder && currentOrder.status === 0 && (
                <Button 
                  key="cancel" 
                  danger 
                  onClick={() => {
                    handleDetailClose();
                    showCancelModal(currentOrder);
                  }}
                >
                  取消订单
                </Button>
              ),
              currentOrder && currentOrder.status === 2 && !currentOrder.is_reviewed && (
                <Button 
                  key="review" 
                  type="primary" 
                  onClick={() => {
                    handleDetailClose();
                    showReviewModal(currentOrder);
                  }}
                >
                  评价服务
                </Button>
              ),
            ].filter(Boolean)}
            width={700}
          >
            {currentOrder ? (
              <>
                <Row gutter={16}>
                  <Col span={24}>
                    {getOrderSteps(currentOrder)}
                  </Col>
                </Row>
                
                <Divider />
                
                <Row gutter={16}>
                  <Col span={8}>
                    <Statistic 
                      title="订单状态" 
                      value={getStatusTag(currentOrder.status)} 
                    />
                  </Col>
                  <Col span={8}>
                    <Statistic 
                      title="订单金额" 
                      value={currentOrder.total_amount} 
                      prefix="¥" 
                    />
                  </Col>
                  <Col span={8}>
                    <Statistic 
                      title="寄养天数" 
                      value={moment(currentOrder.end_date).diff(moment(currentOrder.start_date), 'days') + 1} 
                      suffix="天" 
                    />
                  </Col>
                </Row>
                
                <Divider />
                
                <Descriptions title="基本信息" bordered column={2}>
                  <Descriptions.Item label="订单号">{currentOrder.order_id}</Descriptions.Item>
                  <Descriptions.Item label="创建时间">{moment(currentOrder.created_at).format('YYYY-MM-DD HH:mm:ss')}</Descriptions.Item>
                  <Descriptions.Item label="服务名称">{currentOrder.service_name}</Descriptions.Item>
                  <Descriptions.Item label="服务提供者">{currentOrder.provider_name}</Descriptions.Item>
                  <Descriptions.Item label="宠物名称">{currentOrder.pet_name}</Descriptions.Item>
                  <Descriptions.Item label="宠物类型">
                    {currentOrder.pet_type === 'dog' ? '狗' : 
                     currentOrder.pet_type === 'cat' ? '猫' : 
                     currentOrder.pet_type}
                  </Descriptions.Item>
                  <Descriptions.Item label="开始日期">{moment(currentOrder.start_date).format('YYYY-MM-DD')}</Descriptions.Item>
                  <Descriptions.Item label="结束日期">{moment(currentOrder.end_date).format('YYYY-MM-DD')}</Descriptions.Item>
                  <Descriptions.Item label="备注" span={2}>{currentOrder.remarks || '无'}</Descriptions.Item>
                </Descriptions>
                
                {currentOrder.cancel_reason && (
                  <>
                    <Divider />
                    <Descriptions title="取消信息" bordered>
                      <Descriptions.Item label="取消原因" span={3}>{currentOrder.cancel_reason}</Descriptions.Item>
                      <Descriptions.Item label="取消时间" span={3}>
                        {currentOrder.cancel_time ? moment(currentOrder.cancel_time).format('YYYY-MM-DD HH:mm:ss') : '-'}
                      </Descriptions.Item>
                    </Descriptions>
                  </>
                )}
                
                {currentOrder.review && (
                  <>
                    <Divider />
                    <Card title="我的评价">
                      <div style={{ marginBottom: 8 }}>
                        <Rate disabled value={currentOrder.review.rating} />
                      </div>
                      <Paragraph>{currentOrder.review.content}</Paragraph>
                      <div style={{ fontSize: 12, color: '#999' }}>
                        {moment(currentOrder.review.created_at).format('YYYY-MM-DD HH:mm:ss')}
                      </div>
                    </Card>
                  </>
                )}
              </>
            ) : (
              <div style={{ textAlign: 'center', padding: '50px 0' }}>
                <Spin size="large" />
              </div>
            )}
          </Modal>

          {/* 评价模态框 */}
          <Modal
            title="服务评价"
            visible={reviewVisible}
            onCancel={handleReviewClose}
            footer={[
              <Button key="back" onClick={handleReviewClose}>
                取消
              </Button>,
              <Button 
                key="submit" 
                type="primary" 
                loading={submitting} 
                onClick={handleSubmitReview}
              >
                提交评价
              </Button>,
            ]}
          >
            {currentOrder && (
              <Form
                form={form}
                layout="vertical"
                name="review_form"
              >
                <div style={{ marginBottom: 16 }}>
                  <Card>
                    <Card.Meta
                      title={currentOrder.service_name}
                      description={`订单号: ${currentOrder.order_id} | 宠物: ${currentOrder.pet_name}`}
                    />
                  </Card>
                </div>
                
                <Form.Item
                  name="rating"
                  label="服务评分"
                  rules={[{ required: true, message: '请选择评分' }]}
                  initialValue={5}
                >
                  <Rate />
                </Form.Item>
                
                <Form.Item
                  name="content"
                  label="评价内容"
                  rules={[{ required: true, message: '请输入评价内容' }]}
                >
                  <Input.TextArea 
                    rows={4} 
                    placeholder="请分享您对本次服务的评价和建议"
                  />
                </Form.Item>
              </Form>
            )}
          </Modal>

          {/* 取消订单模态框 */}
          <Modal
            title="取消订单"
            visible={cancelVisible}
            onCancel={handleCancelClose}
            footer={[
              <Button key="back" onClick={handleCancelClose}>
                返回
              </Button>,
              <Button 
                key="submit" 
                type="primary" 
                danger 
                loading={submitting} 
                onClick={handleCancelOrder}
              >
                确认取消
              </Button>,
            ]}
          >
            {currentOrder && (
              <Form
                form={form}
                layout="vertical"
                name="cancel_form"
              >
                <div style={{ marginBottom: 16 }}>
                  <Card>
                    <Card.Meta
                      title={currentOrder.service_name}
                      description={`订单号: ${currentOrder.order_id} | 宠物: ${currentOrder.pet_name}`}
                    />
                  </Card>
                </div>
                
                <Alert
                  message="取消提示"
                  description="取消订单后无法恢复，请确认您的决定。"
                  type="warning"
                  showIcon
                  style={{ marginBottom: 16 }}
                />
                
                <Form.Item
                  name="cancel_reason"
                  label="取消原因"
                  rules={[{ required: true, message: '请输入取消原因' }]}
                >
                  <Input.TextArea 
                    rows={4} 
                    placeholder="请输入取消订单的原因"
                  />
                </Form.Item>
              </Form>
            )}
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default Order;