import React, { useState, useEffect } from 'react';
import { Layout, Menu, Table, Card, Button, Input, Select, Modal, Form, message, Tag, Space, Spin, Popconfirm, InputNumber, Upload } from 'antd';
import { 
  PieChartOutlined, 
  TeamOutlined, 
  ShopOutlined, 
  OrderedListOutlined, 
  LogoutOutlined,
  SearchOutlined,
  EditOutlined,
  DeleteOutlined,
  ExclamationCircleOutlined,
  UploadOutlined,
  PlusOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { api } from '../../utils/api';

const { Header, Sider, Content } = Layout;
const { Option } = Select;
const { TextArea } = Input;

const AdminServiceList = () => {
  const { logout, currentUser } = useAuth();
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const [loading, setLoading] = useState(true);
  const [services, setServices] = useState([]);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });
  const [searchForm] = Form.useForm();
  const [editForm] = Form.useForm();
  const [addForm] = Form.useForm();
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [addModalVisible, setAddModalVisible] = useState(false);
  const [currentService, setCurrentService] = useState(null);
  const [searchParams, setSearchParams] = useState({
    name: '',
    providerId: '',
    status: ''
  });
  const [imageUrl, setImageUrl] = useState('');
  const [fileList, setFileList] = useState([]);
  const [providers, setProviders] = useState([]);

  const fetchServices = async (page = 1, pageSize = 10, params = {}) => {
    try {
      setLoading(true);
      const response = await api.admin.getServices({
        page,
        size: pageSize,
        ...params
      });
      
      if (response.data) {
        setServices(response.data.content || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.totalElements || 0
        });
      }
    } catch (error) {
      console.error('获取服务列表失败:', error);
      message.error('获取服务列表失败');
    } finally {
      setLoading(false);
    }
  };

  const fetchProviders = async () => {
    try {
      const response = await api.admin.getServiceProviders();
      if (response.data) {
        setProviders(response.data || []);
      }
    } catch (error) {
      console.error('获取服务提供商列表失败:', error);
      message.error('获取服务提供商列表失败');
    }
  };

  useEffect(() => {
    fetchServices(pagination.current, pagination.pageSize, searchParams);
    fetchProviders();
  }, []);

  const handleTableChange = (pagination) => {
    fetchServices(pagination.current, pagination.pageSize, searchParams);
  };

  const handleSearch = (values) => {
    const params = {
      name: values.name || '',
      providerId: values.providerId || '',
      status: values.status || ''
    };
    setSearchParams(params);
    fetchServices(1, pagination.pageSize, params);
  };

  const handleReset = () => {
    searchForm.resetFields();
    const params = {
      name: '',
      providerId: '',
      status: ''
    };
    setSearchParams(params);
    fetchServices(1, pagination.pageSize, params);
  };

  const showEditModal = (service) => {
    setCurrentService(service);
    setImageUrl(service.imageUrl);
    setFileList(service.imageUrl ? [{
      uid: '-1',
      name: 'image.png',
      status: 'done',
      url: service.imageUrl,
    }] : []);
    
    editForm.setFieldsValue({
      name: service.name,
      description: service.description,
      price: service.price,
      providerId: service.providerId,
      capacity: service.capacity,
      status: service.status
    });
    setEditModalVisible(true);
  };

  const showAddModal = () => {
    addForm.resetFields();
    setImageUrl('');
    setFileList([]);
    setAddModalVisible(true);
  };

  const handleEditService = async () => {
    try {
      const values = await editForm.validateFields();
      const updatedService = {
        ...values,
        imageUrl
      };
      await api.admin.updateService(currentService.id, updatedService);
      message.success('服务信息更新成功');
      setEditModalVisible(false);
      fetchServices(pagination.current, pagination.pageSize, searchParams);
    } catch (error) {
      console.error('更新服务失败:', error);
      message.error('更新服务失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleAddService = async () => {
    try {
      const values = await addForm.validateFields();
      const newService = {
        ...values,
        imageUrl
      };
      await api.admin.createService(newService);
      message.success('服务创建成功');
      setAddModalVisible(false);
      fetchServices(pagination.current, pagination.pageSize, searchParams);
    } catch (error) {
      console.error('创建服务失败:', error);
      message.error('创建服务失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleDeleteService = async (serviceId) => {
    try {
      await api.admin.deleteService(serviceId);
      message.success('服务删除成功');
      fetchServices(pagination.current, pagination.pageSize, searchParams);
    } catch (error) {
      console.error('删除服务失败:', error);
      message.error('删除服务失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/admin/login');
    message.success('已成功退出登录');
  };

  // 图片上传相关处理
  const beforeUpload = (file) => {
    const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
    if (!isJpgOrPng) {
      message.error('只能上传JPG/PNG格式的图片!');
    }
    const isLt2M = file.size / 1024 / 1024 < 2;
    if (!isLt2M) {
      message.error('图片大小不能超过2MB!');
    }
    return isJpgOrPng && isLt2M;
  };

  const handleUploadChange = (info) => {
    if (info.file.status === 'uploading') {
      return;
    }
    if (info.file.status === 'done') {
      // 假设上传成功后，服务器返回图片URL
      setImageUrl(info.file.response.url);
    }
  };

  // 模拟上传图片
  const customUpload = async ({ file, onSuccess, onError }) => {
    try {
      // 这里应该是实际的图片上传API调用
      // 为了演示，我们模拟一个成功的上传
      setTimeout(() => {
        // 模拟一个图片URL
        const fakeUrl = 'https://example.com/images/' + file.name;
        setImageUrl(fakeUrl);
        onSuccess({ url: fakeUrl });
      }, 1000);
    } catch (error) {
      onError(error);
    }
  };

  const uploadButton = (
    <div>
      <PlusOutlined />
      <div style={{ marginTop: 8 }}>上传</div>
    </div>
  );

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: 80,
    },
    {
      title: '服务名称',
      dataIndex: 'name',
      key: 'name',
      width: 150,
    },
    {
      title: '图片',
      dataIndex: 'imageUrl',
      key: 'imageUrl',
      width: 100,
      render: (imageUrl) => (
        imageUrl ? (
          <img src={imageUrl} alt="服务图片" style={{ width: 50, height: 50, objectFit: 'cover' }} />
        ) : '无图片'
      ),
    },
    {
      title: '描述',
      dataIndex: 'description',
      key: 'description',
      width: 200,
      ellipsis: true,
    },
    {
      title: '价格(元/天)',
      dataIndex: 'price',
      key: 'price',
      width: 120,
      render: (price) => `¥${price.toFixed(2)}`,
    },
    {
      title: '服务提供商',
      dataIndex: 'providerName',
      key: 'providerName',
      width: 150,
    },
    {
      title: '容量',
      dataIndex: 'capacity',
      key: 'capacity',
      width: 100,
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      width: 100,
      render: (status) => {
        let color = status === 'ACTIVE' ? 'green' : 'red';
        let text = status === 'ACTIVE' ? '活跃' : '禁用';
        return <Tag color={color}>{text}</Tag>;
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
      width: 150,
      render: (_, record) => (
        <Space size="small">
          <Button 
            type="primary" 
            icon={<EditOutlined />} 
            size="small" 
            onClick={() => showEditModal(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除此服务吗？"
            icon={<ExclamationCircleOutlined style={{ color: 'red' }} />}
            onConfirm={() => handleDeleteService(record.id)}
            okText="确定"
            cancelText="取消"
          >
            <Button 
              danger 
              icon={<DeleteOutlined />} 
              size="small"
            >
              删除
            </Button>
          </Popconfirm>
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
        <Menu theme="dark" defaultSelectedKeys={['3']} mode="inline">
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
            <h2 style={{ margin: 0 }}>服务管理</h2>
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
            <div style={{ marginBottom: 16, display: 'flex', justifyContent: 'space-between' }}>
              <Form
                form={searchForm}
                layout="inline"
                onFinish={handleSearch}
              >
                <Form.Item name="name">
                  <Input 
                    placeholder="服务名称" 
                    allowClear 
                  />
                </Form.Item>
                <Form.Item name="providerId">
                  <Select 
                    placeholder="服务提供商" 
                    style={{ width: 180 }} 
                    allowClear
                  >
                    {providers.map(provider => (
                      <Option key={provider.id} value={provider.id}>{provider.username}</Option>
                    ))}
                  </Select>
                </Form.Item>
                <Form.Item name="status">
                  <Select 
                    placeholder="状态" 
                    style={{ width: 120 }} 
                    allowClear
                  >
                    <Option value="ACTIVE">活跃</Option>
                    <Option value="INACTIVE">禁用</Option>
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
              <Button 
                type="primary" 
                icon={<PlusOutlined />} 
                onClick={showAddModal}
              >
                添加服务
              </Button>
            </div>

            <Table 
              columns={columns} 
              dataSource={services} 
              rowKey="id" 
              pagination={pagination}
              loading={loading}
              onChange={handleTableChange}
              scroll={{ x: 1300 }}
            />
          </Card>

          {/* 编辑服务模态框 */}
          <Modal
            title="编辑服务"
            visible={editModalVisible}
            onOk={handleEditService}
            onCancel={() => setEditModalVisible(false)}
            okText="保存"
            cancelText="取消"
            width={600}
          >
            <Form
              form={editForm}
              layout="vertical"
            >
              <Form.Item
                name="name"
                label="服务名称"
                rules={[{ required: true, message: '请输入服务名称' }]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="description"
                label="服务描述"
                rules={[{ required: true, message: '请输入服务描述' }]}
              >
                <TextArea rows={4} />
              </Form.Item>
              <Form.Item
                name="price"
                label="价格(元/天)"
                rules={[{ required: true, message: '请输入价格' }]}
              >
                <InputNumber 
                  min={0} 
                  precision={2} 
                  style={{ width: '100%' }} 
                />
              </Form.Item>
              <Form.Item
                name="providerId"
                label="服务提供商"
                rules={[{ required: true, message: '请选择服务提供商' }]}
              >
                <Select>
                  {providers.map(provider => (
                    <Option key={provider.id} value={provider.id}>{provider.username}</Option>
                  ))}
                </Select>
              </Form.Item>
              <Form.Item
                name="capacity"
                label="容量"
                rules={[{ required: true, message: '请输入容量' }]}
              >
                <InputNumber 
                  min={1} 
                  precision={0} 
                  style={{ width: '100%' }} 
                />
              </Form.Item>
              <Form.Item
                name="status"
                label="状态"
                rules={[{ required: true, message: '请选择状态' }]}
              >
                <Select>
                  <Option value="ACTIVE">活跃</Option>
                  <Option value="INACTIVE">禁用</Option>
                </Select>
              </Form.Item>
              <Form.Item
                label="服务图片"
              >
                <Upload
                  name="image"
                  listType="picture-card"
                  className="avatar-uploader"
                  showUploadList={false}
                  beforeUpload={beforeUpload}
                  onChange={handleUploadChange}
                  customRequest={customUpload}
                >
                  {imageUrl ? <img src={imageUrl} alt="服务图片" style={{ width: '100%' }} /> : uploadButton}
                </Upload>
              </Form.Item>
            </Form>
          </Modal>

          {/* 添加服务模态框 */}
          <Modal
            title="添加服务"
            visible={addModalVisible}
            onOk={handleAddService}
            onCancel={() => setAddModalVisible(false)}
            okText="添加"
            cancelText="取消"
            width={600}
          >
            <Form
              form={addForm}
              layout="vertical"
            >
              <Form.Item
                name="name"
                label="服务名称"
                rules={[{ required: true, message: '请输入服务名称' }]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="description"
                label="服务描述"
                rules={[{ required: true, message: '请输入服务描述' }]}
              >
                <TextArea rows={4} />
              </Form.Item>
              <Form.Item
                name="price"
                label="价格(元/天)"
                rules={[{ required: true, message: '请输入价格' }]}
              >
                <InputNumber 
                  min={0} 
                  precision={2} 
                  style={{ width: '100%' }} 
                />
              </Form.Item>
              <Form.Item
                name="providerId"
                label="服务提供商"
                rules={[{ required: true, message: '请选择服务提供商' }]}
              >
                <Select>
                  {providers.map(provider => (
                    <Option key={provider.id} value={provider.id}>{provider.username}</Option>
                  ))}
                </Select>
              </Form.Item>
              <Form.Item
                name="capacity"
                label="容量"
                rules={[{ required: true, message: '请输入容量' }]}
              >
                <InputNumber 
                  min={1} 
                  precision={0} 
                  style={{ width: '100%' }} 
                />
              </Form.Item>
              <Form.Item
                name="status"
                label="状态"
                initialValue="ACTIVE"
                rules={[{ required: true, message: '请选择状态' }]}
              >
                <Select>
                  <Option value="ACTIVE">活跃</Option>
                  <Option value="INACTIVE">禁用</Option>
                </Select>
              </Form.Item>
              <Form.Item
                label="服务图片"
              >
                <Upload
                  name="image"
                  listType="picture-card"
                  className="avatar-uploader"
                  showUploadList={false}
                  beforeUpload={beforeUpload}
                  onChange={handleUploadChange}
                  customRequest={customUpload}
                >
                  {imageUrl ? <img src={imageUrl} alt="服务图片" style={{ width: '100%' }} /> : uploadButton}
                </Upload>
              </Form.Item>
            </Form>
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default AdminServiceList;