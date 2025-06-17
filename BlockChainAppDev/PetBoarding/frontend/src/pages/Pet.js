import React, { useState, useEffect } from 'react';
import { 
  Layout, Menu, Card, Button, Table, Modal, Form, 
  Input, Select, DatePicker, Upload, message, Popconfirm, Spin 
} from 'antd';
import { 
  UserOutlined, PieChartOutlined, ShoppingOutlined, 
  HeartOutlined, BellOutlined, HomeOutlined, 
  PlusOutlined, EditOutlined, DeleteOutlined, UploadOutlined 
} from '@ant-design/icons';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { api } from '../utils/api';
import moment from 'moment';

const { Header, Content, Sider } = Layout;
const { Option } = Select;

const Pet = () => {
  const { user } = useAuth();
  const [form] = Form.useForm();
  const [pets, setPets] = useState([]);
  const [loading, setLoading] = useState(true);
  const [modalVisible, setModalVisible] = useState(false);
  const [modalTitle, setModalTitle] = useState('添加宠物');
  const [editingPet, setEditingPet] = useState(null);
  const [imageUrl, setImageUrl] = useState('');
  const [uploading, setUploading] = useState(false);

  // 获取宠物列表
  const fetchPets = async () => {
    setLoading(true);
    try {
      const response = await api.pet.getPets();
      if (response.data) {
        setPets(response.data);
      }
    } catch (error) {
      console.error('获取宠物列表失败:', error);
      message.error('获取宠物列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchPets();
  }, []);

  // 打开添加宠物模态框
  const showAddModal = () => {
    setModalTitle('添加宠物');
    setEditingPet(null);
    setImageUrl('');
    form.resetFields();
    setModalVisible(true);
  };

  // 打开编辑宠物模态框
  const showEditModal = (pet) => {
    setModalTitle('编辑宠物');
    setEditingPet(pet);
    setImageUrl(pet.image_url || '');
    form.setFieldsValue({
      ...pet,
      birth_date: pet.birth_date ? moment(pet.birth_date) : null,
    });
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
      
      // 转换日期格式
      if (values.birth_date) {
        values.birth_date = values.birth_date.format('YYYY-MM-DD');
      }
      
      // 添加图片URL
      if (imageUrl) {
        values.image_url = imageUrl;
      }

      setUploading(true);
      
      if (editingPet) {
        // 更新宠物
        await api.pet.updatePet(editingPet.pet_id, values);
        message.success('宠物信息更新成功');
      } else {
        // 添加宠物
        await api.pet.addPet(values);
        message.success('宠物添加成功');
      }
      
      setModalVisible(false);
      fetchPets(); // 刷新列表
    } catch (error) {
      console.error('提交宠物信息失败:', error);
      message.error('提交宠物信息失败');
    } finally {
      setUploading(false);
    }
  };

  // 删除宠物
  const handleDelete = async (petId) => {
    try {
      await api.pet.deletePet(petId);
      message.success('宠物删除成功');
      fetchPets(); // 刷新列表
    } catch (error) {
      console.error('删除宠物失败:', error);
      message.error('删除宠物失败');
    }
  };

  // 处理图片上传
  const handleImageUpload = async (info) => {
    if (info.file.status === 'uploading') {
      setUploading(true);
      return;
    }
    
    if (info.file.status === 'done') {
      // 获取上传后的URL
      setImageUrl(info.file.response.url);
      setUploading(false);
      message.success('图片上传成功');
    } else if (info.file.status === 'error') {
      setUploading(false);
      message.error('图片上传失败');
    }
  };

  // 表格列定义
  const columns = [
    {
      title: '宠物照片',
      dataIndex: 'image_url',
      key: 'image_url',
      render: (text) => (
        text ? (
          <img src={text} alt="宠物照片" style={{ width: 50, height: 50, objectFit: 'cover' }} />
        ) : (
          <div style={{ width: 50, height: 50, background: '#eee', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <HomeOutlined />
          </div>
        )
      ),
    },
    {
      title: '宠物名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '宠物类型',
      dataIndex: 'type',
      key: 'type',
      render: (text) => {
        const typeMap = {
          'dog': '狗',
          'cat': '猫',
          'bird': '鸟',
          'fish': '鱼',
          'other': '其他'
        };
        return typeMap[text] || text;
      },
    },
    {
      title: '品种',
      dataIndex: 'breed',
      key: 'breed',
    },
    {
      title: '年龄',
      dataIndex: 'age',
      key: 'age',
      render: (text) => `${text} 岁`,
    },
    {
      title: '性别',
      dataIndex: 'gender',
      key: 'gender',
      render: (text) => text === 'male' ? '公' : '母',
    },
    {
      title: '操作',
      key: 'action',
      render: (_, record) => (
        <>
          <Button 
            type="link" 
            icon={<EditOutlined />} 
            onClick={() => showEditModal(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确定要删除这个宠物吗？"
            onConfirm={() => handleDelete(record.pet_id)}
            okText="确定"
            cancelText="取消"
          >
            <Button 
              type="link" 
              danger 
              icon={<DeleteOutlined />}
            >
              删除
            </Button>
          </Popconfirm>
        </>
      ),
    },
  ];

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider width={200} className="site-layout-background">
        <Menu
          mode="inline"
          defaultSelectedKeys={['pets']}
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
            title="我的宠物"
            extra={
              <Button 
                type="primary" 
                icon={<PlusOutlined />} 
                onClick={showAddModal}
              >
                添加宠物
              </Button>
            }
          >
            {loading ? (
              <div style={{ textAlign: 'center', padding: '50px 0' }}>
                <Spin size="large" />
              </div>
            ) : (
              <Table 
                columns={columns} 
                dataSource={pets} 
                rowKey="pet_id" 
                pagination={{ pageSize: 10 }}
                locale={{ emptyText: '暂无宠物信息' }}
              />
            )}
          </Card>

          {/* 添加/编辑宠物模态框 */}
          <Modal
            title={modalTitle}
            visible={modalVisible}
            onCancel={handleCancel}
            footer={[
              <Button key="back" onClick={handleCancel}>
                取消
              </Button>,
              <Button 
                key="submit" 
                type="primary" 
                loading={uploading} 
                onClick={handleSubmit}
              >
                提交
              </Button>,
            ]}
          >
            <Form
              form={form}
              layout="vertical"
              name="pet_form"
            >
              <Form.Item
                name="name"
                label="宠物名称"
                rules={[{ required: true, message: '请输入宠物名称' }]}
              >
                <Input placeholder="请输入宠物名称" />
              </Form.Item>

              <Form.Item
                name="type"
                label="宠物类型"
                rules={[{ required: true, message: '请选择宠物类型' }]}
              >
                <Select placeholder="请选择宠物类型">
                  <Option value="dog">狗</Option>
                  <Option value="cat">猫</Option>
                  <Option value="bird">鸟</Option>
                  <Option value="fish">鱼</Option>
                  <Option value="other">其他</Option>
                </Select>
              </Form.Item>

              <Form.Item
                name="breed"
                label="品种"
                rules={[{ required: true, message: '请输入宠物品种' }]}
              >
                <Input placeholder="请输入宠物品种" />
              </Form.Item>

              <Form.Item
                name="age"
                label="年龄"
                rules={[{ required: true, message: '请输入宠物年龄' }]}
              >
                <Input type="number" min={0} placeholder="请输入宠物年龄" />
              </Form.Item>

              <Form.Item
                name="gender"
                label="性别"
                rules={[{ required: true, message: '请选择宠物性别' }]}
              >
                <Select placeholder="请选择宠物性别">
                  <Option value="male">公</Option>
                  <Option value="female">母</Option>
                </Select>
              </Form.Item>

              <Form.Item
                name="birth_date"
                label="出生日期"
              >
                <DatePicker style={{ width: '100%' }} />
              </Form.Item>

              <Form.Item
                name="weight"
                label="体重(kg)"
              >
                <Input type="number" min={0} step={0.1} placeholder="请输入宠物体重" />
              </Form.Item>

              <Form.Item
                name="health_info"
                label="健康信息"
              >
                <Input.TextArea rows={4} placeholder="请输入宠物健康信息，如疫苗接种情况、过敏史等" />
              </Form.Item>

              <Form.Item
                name="special_needs"
                label="特殊需求"
              >
                <Input.TextArea rows={4} placeholder="请输入宠物的特殊需求，如饮食习惯、行为特点等" />
              </Form.Item>

              <Form.Item
                label="宠物照片"
              >
                <Upload
                  name="file"
                  listType="picture-card"
                  className="avatar-uploader"
                  showUploadList={false}
                  action="/api/upload"
                  onChange={handleImageUpload}
                >
                  {imageUrl ? (
                    <img src={imageUrl} alt="宠物照片" style={{ width: '100%' }} />
                  ) : (
                    <div>
                      {uploading ? <Spin /> : <UploadOutlined />}
                      <div style={{ marginTop: 8 }}>上传照片</div>
                    </div>
                  )}
                </Upload>
              </Form.Item>
            </Form>
          </Modal>
        </Content>
      </Layout>
    </Layout>
  );
};

export default Pet;