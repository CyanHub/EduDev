import React, { useState, useEffect } from 'react';
import { 
  Layout, Menu, Card, Button, List, Rate, 
  Modal, Form, Input, Divider, message, Spin, Typography, Avatar 
} from 'antd';
import { 
  UserOutlined, PieChartOutlined, ShoppingOutlined, 
  HeartOutlined, BellOutlined, HomeOutlined, 
  EditOutlined, DeleteOutlined, ExclamationCircleOutlined 
} from '@ant-design/icons';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { TextArea } = Input;
const { Title, Text, Paragraph } = Typography;
const { confirm } = Modal;

const Review = () => {
  const { user } = useAuth();
  const [form] = Form.useForm();
  const [reviews, setReviews] = useState([]);
  const [loading, setLoading] = useState(true);
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [currentReview, setCurrentReview] = useState(null);
  const [submitting, setSubmitting] = useState(false);

  // 获取评价列表
  const fetchReviews = async () => {
    setLoading(true);
    try {
      const response = await api.review.getUserReviews();
      if (response.data) {
        setReviews(response.data);
      }
    } catch (error) {
      console.error('获取评价列表失败:', error);
      message.error('获取评价列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchReviews();
  }, []);

  // 打开编辑评价模态框
  const showEditModal = (review) => {
    setCurrentReview(review);
    form.setFieldsValue({
      rating: review.rating,
      content: review.content,
    });
    setEditModalVisible(true);
  };

  // 关闭编辑评价模态框
  const handleEditCancel = () => {
    setEditModalVisible(false);
  };

  // 提交编辑评价
  const handleEditSubmit = async () => {
    try {
      const values = await form.validateFields();
      setSubmitting(true);
      
      await api.review.updateReview(currentReview.review_id, {
        rating: values.rating,
        content: values.content,
      });
      
      message.success('评价更新成功');
      setEditModalVisible(false);
      fetchReviews();
    } catch (error) {
      console.error('更新评价失败:', error);
      message.error('更新评价失败');
    } finally {
      setSubmitting(false);
    }
  };

  // 删除评价
  const handleDelete = (review) => {
    confirm({
      title: '确认删除',
      icon: <ExclamationCircleOutlined />,
      content: '确定要删除这条评价吗？删除后无法恢复。',
      onOk: async () => {
        try {
          await api.review.deleteReview(review.review_id);
          message.success('评价删除成功');
          fetchReviews();
        } catch (error) {
          console.error('删除评价失败:', error);
          message.error('删除评价失败');
        }
      },
    });
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['reviews']}
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
          <Card title="我的评价">
            {loading ? (
              <div style={{ textAlign: 'center', padding: '50px 0' }}>
                <Spin size="large" />
              </div>
            ) : (
              <List
                itemLayout="vertical"
                dataSource={reviews}
                locale={{ emptyText: '暂无评价' }}
                pagination={{
                  pageSize: 10,
                  showTotal: (total) => `共 ${total} 条评价`,
                }}
                renderItem={item => (
                  <List.Item
                    key={item.review_id}
                    actions={[
                      <Button 
                        type="link" 
                        icon={<EditOutlined />} 
                        onClick={() => showEditModal(item)}
                      >
                        编辑
                      </Button>,
                      <Button 
                        type="link" 
                        danger 
                        icon={<DeleteOutlined />} 
                        onClick={() => handleDelete(item)}
                      >
                        删除
                      </Button>,
                    ]}
                  >
                    <List.Item.Meta
                      avatar={<Avatar icon={<HomeOutlined />} />}
                      title={
                        <div>
                          <Link to={`/services/${item.service_id}`}>{item.service_name}</Link>
                          <div style={{ float: 'right' }}>
                            <Rate disabled defaultValue={item.rating} />
                          </div>
                        </div>
                      }
                      description={
                        <>
                          <div>订单号: <Link to={`/orders/${item.order_id}`}>{item.order_id}</Link></div>
                          <div>评价时间: {moment(item.created_at).format('YYYY-MM-DD HH:mm:ss')}</div>
                        </>
                      }
                    />
                    <div style={{ marginTop: 16 }}>
                      <Paragraph>{item.content}</Paragraph>
                    </div>
                    {item.reply && (
                      <div style={{ background: '#f5f5f5', padding: 16, marginTop: 16, borderRadius: 4 }}>
                        <Text strong>商家回复：</Text>
                        <Paragraph>{item.reply}</Paragraph>
                        <div style={{ fontSize: 12, color: '#999' }}>
                          {moment(item.reply_time).format('YYYY-MM-DD HH:mm:ss')}
                        </div>
                      </div>
                    )}
                  </List.Item>
                )}
              />
            )}
          </Card>

          {/* 编辑评价模态框 */}
          <Modal
            title="编辑评价"
            visible={editModalVisible}
            onCancel={handleEditCancel}
            footer={[
              <Button key="back" onClick={handleEditCancel}>
                取消
              </Button>,
              <Button 
                key="submit" 
                type="primary" 
                loading={submitting} 
                onClick={handleEditSubmit}
              >
                提交
              </Button>,
            ]}
          >
            {currentReview && (
              <Form
                form={form}
                layout="vertical"
                name="edit_review_form"
                initialValues={{
                  rating: currentReview.rating,
                  content: currentReview.content,
                }}
              >
                <div style={{ marginBottom: 16 }}>
                  <Card>
                    <Card.Meta
                      title={currentReview.service_name}
                      description={`订单号: ${currentReview.order_id}`}
                    />
                  </Card>
                </div>
                
                <Form.Item
                  name="rating"
                  label="服务评分"
                  rules={[{ required: true, message: '请选择评分' }]}
                >
                  <Rate />
                </Form.Item>
                
                <Form.Item
                  name="content"
                  label="评价内容"
                  rules={[{ required: true, message: '请输入评价内容' }]}
                >
                  <TextArea 
                    rows={4} 
                    placeholder="请分享您对本次服务的评价和建议"
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

export default Review;