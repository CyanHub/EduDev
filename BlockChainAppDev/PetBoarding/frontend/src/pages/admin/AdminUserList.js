import React, { useState, useEffect } from 'react';
import { Layout, Menu, Table, Card, Button, Input, Select, Modal, Form, message, Tag, Space, Spin, Popconfirm } from 'antd';
import { 
  UserOutlined, 
  PieChartOutlined, 
  TeamOutlined, 
  ShopOutlined, 
  OrderedListOutlined, 
  LogoutOutlined,
  SearchOutlined,
  EditOutlined,
  DeleteOutlined,
  ExclamationCircleOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { api } from '../../utils/api';

const { Header, Sider, Content } = Layout;
const { Option } = Select;

const AdminUserList = () => {
  const { logout, currentUser } = useAuth();
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const [loading, setLoading] = useState(true);
  const [users, setUsers] = useState([]);
  const [pagination, setPagination] = useState({
    current: 1,
    pageSize: 10,
    total: 0
  });
  const [searchForm] = Form.useForm();
  const [editForm] = Form.useForm();
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [currentUser2Edit, setCurrentUser2Edit] = useState(null);
  const [searchParams, setSearchParams] = useState({
    username: '',
    email: '',
    role: ''
  });

  const fetchUsers = async (page = 1, pageSize = 10, params = {}) => {
    try {
      setLoading(true);
      const response = await api.admin.getUsers({
        page,
        size: pageSize,
        ...params
      });
      
      if (response.data) {
        setUsers(response.data.content || []);
        setPagination({
          current: page,
          pageSize,
          total: response.data.totalElements || 0
        });
      }
    } catch (error) {
      console.error('获取用户列表失败:', error);
      message.error('获取用户列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers(pagination.current, pagination.pageSize, searchParams);
  }, []);

  const handleTableChange = (pagination) => {
    fetchUsers(pagination.current, pagination.pageSize, searchParams);
  };

  const handleSearch = (values) => {
    const params = {
      username: values.username || '',
      email: values.email || '',
      role: values.role || ''
    };
    setSearchParams(params);
    fetchUsers(1, pagination.pageSize, params);
  };

  const handleReset = () => {
    searchForm.resetFields();
    const params = {
      username: '',
      email: '',
      role: ''
    };
    setSearchParams(params);
    fetchUsers(1, pagination.pageSize, params);
  };

  const showEditModal = (user) => {
    setCurrentUser2Edit(user);
    editForm.setFieldsValue({
      username: user.username,
      email: user.email,
      phone: user.phone,
      role: user.role,
      status: user.status
    });
    setEditModalVisible(true);
  };

  const handleEditUser = async () => {
    try {
      const values = await editForm.validateFields();
      await api.admin.updateUser(currentUser2Edit.id, values);
      message.success('用户信息更新成功');
      setEditModalVisible(false);
      fetchUsers(pagination.current, pagination.pageSize, searchParams);
    } catch (error) {
      console.error('更新用户失败:', error);
      message.error('更新用户失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleDeleteUser = async (userId) => {
    try {
      await api.admin.deleteUser(userId);
      message.success('用户删除成功');
      fetchUsers(pagination.current, pagination.pageSize, searchParams);
    } catch (error) {
      console.error('删除用户失败:', error);
      message.error('删除用户失败: ' + (error.response?.data?.message || error.message));
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/admin/login');
    message.success('已成功退出登录');
  };

  const columns = [
    {
      title: 'ID',
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
      width: 120,
      render: (role) => {
        let color = role === 'SERVICE_PROVIDER' ? 'green' : 'blue';
        let text = role === 'SERVICE_PROVIDER' ? '服务提供商' : '普通用户';
        return <Tag color={color}>{text}</Tag>;
      },
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
      title: '注册时间',
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
            title="确定要删除此用户吗？"
            icon={<ExclamationCircleOutlined style={{ color: 'red' }} />}
            onConfirm={() => handleDeleteUser(record.id)}
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
        <Menu theme="dark" defaultSelectedKeys={['2']} mode="inline">
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
            <h2 style={{ margin: 0 }}>用户管理</h2>
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
              <Form.Item name="username">
                <Input 
                  placeholder="用户名" 
                  prefix={<UserOutlined />} 
                  allowClear 
                />
              </Form.Item>
              <Form.Item name="email">
                <Input 
                  placeholder="邮箱" 
                  allowClear 
                />
              </Form.Item>
              <Form.Item name="role">
                <Select 
                  placeholder="角色" 
                  style={{ width: 120 }} 
                  allowClear
                >
                  <Option value="USER">普通用户</Option>
                  <Option value="SERVICE_PROVIDER">服务提供商</Option>
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
              dataSource={users} 
              rowKey="id" 
              pagination={pagination}
              loading={loading}
              onChange={handleTableChange}
              scroll={{ x: 1100 }}
            />
          </Card>

          <Modal
            title="编辑用户"
            visible={editModalVisible}
            onOk={handleEditUser}
            onCancel={() => setEditModalVisible(false)}
            okText="保存"
            cancelText="取消"
          >
            <Form
              form={editForm}
              layout="vertical"
            >
              <Form.Item
                name="username"
                label="用户名"
                rules={[{ required: true, message: '请输入用户名' }]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="email"
                label="邮箱"
                rules={[
                  { required: true, message: '请输入邮箱' },
                  { type: 'email', message: '请输入有效的邮箱地址' }
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="phone"
                label="电话"
                rules={[{ required: true, message: '请输入电话号码' }]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                name="role"
                label="角色"
                rules={[{ required: true, message: '请选择角色' }]}
              >
                <Select>
                  <Option value="USER">普通用户</Option>
                  <Option value="SERVICE_PROVIDER">服务提供商</Option>
                </Select>
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
            </Form>
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default AdminUserList;