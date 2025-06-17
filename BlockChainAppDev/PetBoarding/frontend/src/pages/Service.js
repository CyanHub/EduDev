import React, { useState, useEffect } from 'react';
import { 
  Layout, Menu, Card, Button, Row, Col, Input, Select, 
  DatePicker, Modal, Form, InputNumber, Rate, Tag, 
  Pagination, Spin, message, Divider, List, Avatar, Typography 
} from 'antd';
import { 
  UserOutlined, PieChartOutlined, ShoppingOutlined, 
  HeartOutlined, BellOutlined, HomeOutlined, 
  SearchOutlined, FilterOutlined, EnvironmentOutlined,
  ClockCircleOutlined, DollarOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { Option } = Select;
const { RangePicker } = DatePicker;
const { Title, Text, Paragraph } = Typography;

const Service = () => {
  const { user } = useAuth();
  const navigate = useNavigate();
  const [form] = Form.useForm();
  const [services, setServices] = useState([]);
  const [loading, setLoading] = useState(true);
  const [total, setTotal] = useState(0);
  const [currentPage, setCurrentPage] = useState(1);
  const [pageSize, setPageSize] = useState(8);
  const [modalVisible, setModalVisible] = useState(false);
  const [selectedService, setSelectedService] = useState(null);
  const [selectedPet, setSelectedPet] = useState(null);
  const [pets, setPets] = useState([]);
  const [searchParams, setSearchParams] = useState({
    keyword: '',
    type: '',
    min_price: '',
    max_price: '',
    min_rating: '',
    location: '',
  });
  const [bookingDates, setBookingDates] = useState([]);
  const [submitting, setSubmitting] = useState(false);

  // 获取寄养服务列表
  const fetchServices = async (page = 1, size = 8, params = {}) => {
    setLoading(true);
    try {
      const response = await api.service.getServices({
        page,
        size,
        ...params
      });
      if (response.data) {
        setServices(response.data.items || []);
        setTotal(response.data.total || 0);
      }
    } catch (error) {
      console.error('获取寄养服务列表失败:', error);
      message.error('获取寄养服务列表失败');
    } finally {
      setLoading(false);
    }
  };

  // 获取用户宠物列表
  const fetchPets = async () => {
    try {
      const response = await api.pet.getPets();
      if (response.data) {
        setPets(response.data);
      }
    } catch (error) {
      console.error('获取宠物列表失败:', error);
    }
  };

  useEffect(() => {
    fetchServices(currentPage, pageSize, searchParams);
    fetchPets();
  }, [currentPage, pageSize]);

  // 处理搜索
  const handleSearch = () => {
    setCurrentPage(1);
    fetchServices(1, pageSize, searchParams);
  };

  // 重置搜索
  const handleReset = () => {
    setSearchParams({
      keyword: '',
      type: '',
      min_price: '',
      max_price: '',
      min_rating: '',
      location: '',
    });
    setCurrentPage(1);
    fetchServices(1, pageSize, {});
  };

  // 处理分页变化
  const handlePageChange = (page, size) => {
    setCurrentPage(page);
    setPageSize(size);
  };

  // 打开预订模态框
  const showBookingModal = (service) => {
    if (!user) {
      message.warning('请先登录');
      navigate('/login');
      return;
    }
    
    if (pets.length === 0) {
      message.warning('请先添加宠物');
      navigate('/pets');
      return;
    }
    
    setSelectedService(service);
    setSelectedPet(null);
    setBookingDates([]);
    form.resetFields();
    setModalVisible(true);
  };

  // 关闭模态框
  const handleCancel = () => {
    setModalVisible(false);
  };

  // 处理表单提交
  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();
      
      if (!selectedPet) {
        message.error('请选择宠物');
        return;
      }
      
      if (!bookingDates || bookingDates.length !== 2) {
        message.error('请选择寄养日期');
        return;
      }
      
      const startDate = bookingDates[0].format('YYYY-MM-DD');
      const endDate = bookingDates[1].format('YYYY-MM-DD');
      
      // 计算天数
      const days = bookingDates[1].diff(bookingDates[0], 'days') + 1;
      
      // 计算总价
      const totalAmount = days * selectedService.price;
      
      const orderData = {
        service_id: selectedService.service_id,
        pet_id: selectedPet.pet_id,
        start_date: startDate,
        end_date: endDate,
        total_amount: totalAmount,
        remarks: values.remarks || '',
      };
      
      setSubmitting(true);
      
      // 创建订单
      const response = await api.order.createOrder(orderData);
      
      if (response.data) {
        message.success('预订成功');
        setModalVisible(false);
        navigate(`/orders/${response.data.order_id}`);
      }
    } catch (error) {
      console.error('提交订单失败:', error);
      message.error('提交订单失败');
    } finally {
      setSubmitting(false);
    }
  };

  // 处理日期选择
  const handleDateChange = (dates) => {
    setBookingDates(dates);
  };

  // 处理宠物选择
  const handlePetSelect = (petId) => {
    const pet = pets.find(p => p.pet_id === petId);
    setSelectedPet(pet);
  };

  // 渲染服务卡片
  const renderServiceCard = (service) => {
    return (
      <Card
        hoverable
        className="service-card"
        cover={
          service.image_url ? (
            <img 
              alt={service.name} 
              src={service.image_url} 
              style={{ height: 200, objectFit: 'cover' }} 
            />
          ) : (
            <div style={{ height: 200, background: '#f5f5f5', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
              <HomeOutlined style={{ fontSize: 48, color: '#d9d9d9' }} />
            </div>
          )
        }
      >
        <div style={{ marginBottom: 8 }}>
          <Tag color="blue">{service.type === 'dog' ? '狗' : service.type === 'cat' ? '猫' : service.type}</Tag>
          <Rate disabled defaultValue={service.rating || 0} style={{ fontSize: 12, marginLeft: 8 }} />
        </div>
        
        <Card.Meta
          title={service.name}
          description={
            <>
              <div style={{ display: 'flex', alignItems: 'center', marginBottom: 4 }}>
                <EnvironmentOutlined style={{ marginRight: 4 }} />
                <Text type="secondary">{service.location || '未知地点'}</Text>
              </div>
              <div style={{ display: 'flex', alignItems: 'center', marginBottom: 4 }}>
                <DollarOutlined style={{ marginRight: 4 }} />
                <Text type="danger">¥{service.price}/天</Text>
              </div>
              <Paragraph ellipsis={{ rows: 2 }}>
                {service.description || '暂无描述'}
              </Paragraph>
            </>
          }
        />
        
        <div style={{ marginTop: 16, textAlign: 'right' }}>
          <Button 
            type="primary" 
            onClick={() => showBookingModal(service)}
          >
            立即预订
          </Button>
        </div>
      </Card>
    );
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['services']}
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
          <Card style={{ marginBottom: 24 }}>
            <Form layout="inline" style={{ marginBottom: 16 }}>
              <Form.Item>
                <Input
                  placeholder="搜索服务名称"
                  value={searchParams.keyword}
                  onChange={(e) => setSearchParams({ ...searchParams, keyword: e.target.value })}
                  prefix={<SearchOutlined />}
                  style={{ width: 200 }}
                />
              </Form.Item>
              
              <Form.Item>
                <Select
                  placeholder="宠物类型"
                  value={searchParams.type}
                  onChange={(value) => setSearchParams({ ...searchParams, type: value })}
                  style={{ width: 120 }}
                  allowClear
                >
                  <Option value="dog">狗</Option>
                  <Option value="cat">猫</Option>
                  <Option value="bird">鸟</Option>
                  <Option value="other">其他</Option>
                </Select>
              </Form.Item>
              
              <Form.Item>
                <Select
                  placeholder="价格区间"
                  onChange={(value) => {
                    if (value) {
                      const [min, max] = value.split('-');
                      setSearchParams({ 
                        ...searchParams, 
                        min_price: min, 
                        max_price: max === 'max' ? '' : max 
                      });
                    } else {
                      setSearchParams({ 
                        ...searchParams, 
                        min_price: '', 
                        max_price: '' 
                      });
                    }
                  }}
                  style={{ width: 120 }}
                  allowClear
                >
                  <Option value="0-100">¥0-100/天</Option>
                  <Option value="100-200">¥100-200/天</Option>
                  <Option value="200-300">¥200-300/天</Option>
                  <Option value="300-max">¥300以上/天</Option>
                </Select>
              </Form.Item>
              
              <Form.Item>
                <Select
                  placeholder="最低评分"
                  value={searchParams.min_rating}
                  onChange={(value) => setSearchParams({ ...searchParams, min_rating: value })}
                  style={{ width: 120 }}
                  allowClear
                >
                  <Option value="3">3星及以上</Option>
                  <Option value="4">4星及以上</Option>
                  <Option value="4.5">4.5星及以上</Option>
                </Select>
              </Form.Item>
              
              <Form.Item>
                <Button 
                  type="primary" 
                  icon={<SearchOutlined />} 
                  onClick={handleSearch}
                >
                  搜索
                </Button>
              </Form.Item>
              
              <Form.Item>
                <Button onClick={handleReset}>重置</Button>
              </Form.Item>
            </Form>
          </Card>
          
          {loading ? (
            <div style={{ textAlign: 'center', padding: '50px 0' }}>
              <Spin size="large" />
            </div>
          ) : (
            <>
              <Row gutter={[16, 16]}>
                {services.length > 0 ? (
                  services.map(service => (
                    <Col xs={24} sm={12} md={8} lg={6} key={service.service_id}>
                      {renderServiceCard(service)}
                    </Col>
                  ))
                ) : (
                  <Col span={24}>
                    <div style={{ textAlign: 'center', padding: '50px 0' }}>
                      <Text type="secondary">暂无寄养服务</Text>
                    </div>
                  </Col>
                )}
              </Row>
              
              {total > 0 && (
                <div style={{ textAlign: 'right', marginTop: 16 }}>
                  <Pagination
                    current={currentPage}
                    pageSize={pageSize}
                    total={total}
                    onChange={handlePageChange}
                    showSizeChanger
                    showQuickJumper
                    showTotal={(total) => `共 ${total} 条`}
                  />
                </div>
              )}
            </>
          )}
          
          {/* 预订模态框 */}
          <Modal
            title="预订寄养服务"
            visible={modalVisible}
            onCancel={handleCancel}
            footer={[
              <Button key="back" onClick={handleCancel}>
                取消
              </Button>,
              <Button 
                key="submit" 
                type="primary" 
                loading={submitting} 
                onClick={handleSubmit}
              >
                提交订单
              </Button>,
            ]}
            width={600}
          >
            {selectedService && (
              <>
                <div style={{ marginBottom: 16 }}>
                  <Card>
                    <Card.Meta
                      avatar={selectedService.image_url ? (
                        <Avatar src={selectedService.image_url} size={64} />
                      ) : (
                        <Avatar icon={<HomeOutlined />} size={64} />
                      )}
                      title={selectedService.name}
                      description={
                        <>
                          <div>价格: ¥{selectedService.price}/天</div>
                          <div>类型: {selectedService.type === 'dog' ? '狗' : selectedService.type === 'cat' ? '猫' : selectedService.type}</div>
                          <div>地点: {selectedService.location || '未知'}</div>
                        </>
                      }
                    />
                  </Card>
                </div>
                
                <Form
                  form={form}
                  layout="vertical"
                  name="booking_form"
                >
                  <Form.Item
                    name="pet_id"
                    label="选择宠物"
                    rules={[{ required: true, message: '请选择宠物' }]}
                  >
                    <Select 
                      placeholder="请选择宠物"
                      onChange={handlePetSelect}
                    >
                      {pets.map(pet => (
                        <Option key={pet.pet_id} value={pet.pet_id}>
                          {pet.name} ({pet.type === 'dog' ? '狗' : pet.type === 'cat' ? '猫' : pet.type})
                        </Option>
                      ))}
                    </Select>
                  </Form.Item>
                  
                  <Form.Item
                    name="dates"
                    label="寄养日期"
                    rules={[{ required: true, message: '请选择寄养日期' }]}
                  >
                    <RangePicker 
                      style={{ width: '100%' }}
                      disabledDate={(current) => current && current < moment().startOf('day')}
                      onChange={handleDateChange}
                    />
                  </Form.Item>
                  
                  {bookingDates && bookingDates.length === 2 && (
                    <div style={{ marginBottom: 16 }}>
                      <Card>
                        <div>寄养天数: {bookingDates[1].diff(bookingDates[0], 'days') + 1} 天</div>
                        <div>总价: ¥{(bookingDates[1].diff(bookingDates[0], 'days') + 1) * selectedService.price}</div>
                      </Card>
                    </div>
                  )}
                  
                  <Form.Item
                    name="remarks"
                    label="备注"
                  >
                    <Input.TextArea 
                      rows={4} 
                      placeholder="请输入特殊要求或备注信息"
                    />
                  </Form.Item>
                </Form>
              </>
            )}
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default Service;