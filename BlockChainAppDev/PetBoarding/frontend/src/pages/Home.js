import React from 'react';
import { Layout, Typography, Button, Row, Col, Card, Carousel, Statistic, Divider, List, Avatar } from 'antd';
import { 
  HeartOutlined, 
  HomeOutlined, 
  SafetyOutlined, 
  CustomerServiceOutlined,
  StarOutlined,
  UserOutlined,
  ShopOutlined,
  CommentOutlined
} from '@ant-design/icons';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const { Header, Content, Footer } = Layout;
const { Title, Paragraph, Text } = Typography;

const Home = () => {
  const { currentUser } = useAuth();
  const navigate = useNavigate();

  // 模拟数据
  const featuredServices = [
    {
      id: 1,
      name: '豪华寄养套餐',
      description: '为您的爱宠提供豪华的住宿环境，包含日常喂食、遛狗、洗澡和基础健康检查。',
      price: 199,
      imageUrl: 'https://via.placeholder.com/300x200?text=Premium+Pet+Boarding',
      rating: 4.9,
      provider: '爱宠之家',
    },
    {
      id: 2,
      name: '标准寄养服务',
      description: '提供舒适的住宿环境，包含日常喂食、定时遛狗和基础照顾。',
      price: 129,
      imageUrl: 'https://via.placeholder.com/300x200?text=Standard+Pet+Boarding',
      rating: 4.7,
      provider: '宠物乐园',
    },
    {
      id: 3,
      name: '猫咪专属寄养',
      description: '为猫咪提供安静、舒适的环境，配备猫爬架、猫砂盆和专业猫咪照顾。',
      price: 159,
      imageUrl: 'https://via.placeholder.com/300x200?text=Cat+Boarding',
      rating: 4.8,
      provider: '喵星人之家',
    },
    {
      id: 4,
      name: '小型宠物寄养',
      description: '适合兔子、仓鼠等小型宠物的专业寄养服务，提供专业的饲养和照顾。',
      price: 99,
      imageUrl: 'https://via.placeholder.com/300x200?text=Small+Pet+Boarding',
      rating: 4.6,
      provider: '小宠之家',
    },
  ];

  const testimonials = [
    {
      id: 1,
      name: '张先生',
      avatar: 'https://via.placeholder.com/50',
      content: '我的金毛在这里寄养了一周，回来后状态非常好，看得出来受到了很好的照顾。工作人员也很专业，会定期发送照片和视频，让我随时了解狗狗的情况。',
      rating: 5,
    },
    {
      id: 2,
      name: '李女士',
      avatar: 'https://via.placeholder.com/50',
      content: '猫咪很挑剔，但在这里住得很开心。环境干净，工作人员很有耐心，会按照我的要求照顾它。价格合理，服务周到，非常满意。',
      rating: 5,
    },
    {
      id: 3,
      name: '王先生',
      avatar: 'https://via.placeholder.com/50',
      content: '第一次把宠物寄养，很担心，但这里的服务让我很放心。环境很好，工作人员很专业，会定时发送宠物的照片和视频。下次还会选择这里。',
      rating: 4,
    },
  ];

  return (
    <Layout className="layout">
      <Header style={{ 
        position: 'fixed', 
        zIndex: 1, 
        width: '100%', 
        background: '#fff', 
        boxShadow: '0 2px 8px rgba(0,0,0,0.06)',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '0 50px'
      }}>
        <div className="logo" style={{ fontSize: '24px', fontWeight: 'bold' }}>
          宠物寄养系统
        </div>
        <div>
          {currentUser ? (
            <Button type="primary" onClick={() => navigate('/dashboard')}>
              进入控制台
            </Button>
          ) : (
            <>
              <Button type="primary" style={{ marginRight: 16 }} onClick={() => navigate('/login')}>
                登录
              </Button>
              <Button onClick={() => navigate('/register')}>
                注册
              </Button>
            </>
          )}
        </div>
      </Header>
      <Content style={{ padding: '0 50px', marginTop: 64 }}>
        {/* 轮播图 */}
        <div style={{ background: '#fff', padding: '24px 0' }}>
          <Carousel autoplay>
            <div>
              <div style={{ 
                height: '500px', 
                background: 'url(https://via.placeholder.com/1200x500?text=Pet+Boarding+Service) center/cover no-repeat',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                color: '#fff',
                textShadow: '0 0 10px rgba(0,0,0,0.5)'
              }}>
                <Title style={{ color: '#fff' }}>为您的爱宠提供最贴心的照顾</Title>
                <Paragraph style={{ color: '#fff', fontSize: '18px' }}>
                  专业的宠物寄养服务，让您的宠物在您外出时也能得到最好的照顾
                </Paragraph>
                <Button type="primary" size="large" onClick={() => navigate('/service')}>
                  立即预订
                </Button>
              </div>
            </div>
            <div>
              <div style={{ 
                height: '500px', 
                background: 'url(https://via.placeholder.com/1200x500?text=Professional+Care) center/cover no-repeat',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                color: '#fff',
                textShadow: '0 0 10px rgba(0,0,0,0.5)'
              }}>
                <Title style={{ color: '#fff' }}>专业的宠物护理团队</Title>
                <Paragraph style={{ color: '#fff', fontSize: '18px' }}>
                  我们的团队由经验丰富的宠物护理专家组成，为您的宠物提供全方位的照顾
                </Paragraph>
                <Button type="primary" size="large" onClick={() => navigate('/service')}>
                  了解更多
                </Button>
              </div>
            </div>
            <div>
              <div style={{ 
                height: '500px', 
                background: 'url(https://via.placeholder.com/1200x500?text=Comfortable+Environment) center/cover no-repeat',
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                color: '#fff',
                textShadow: '0 0 10px rgba(0,0,0,0.5)'
              }}>
                <Title style={{ color: '#fff' }}>舒适安全的寄养环境</Title>
                <Paragraph style={{ color: '#fff', fontSize: '18px' }}>
                  宽敞明亮的活动空间，干净整洁的住宿环境，让您的宠物感到宾至如归
                </Paragraph>
                <Button type="primary" size="large" onClick={() => navigate('/service')}>
                  查看环境
                </Button>
              </div>
            </div>
          </Carousel>
        </div>

        {/* 服务特点 */}
        <div style={{ background: '#fff', padding: '50px 0' }}>
          <Title level={2} style={{ textAlign: 'center', marginBottom: 50 }}>我们的服务特点</Title>
          <Row gutter={[32, 32]} justify="center">
            <Col xs={24} sm={12} md={6}>
              <Card className="feature-card" bordered={false} style={{ textAlign: 'center' }}>
                <HeartOutlined style={{ fontSize: 48, color: '#ff4d4f', marginBottom: 16 }} />
                <Title level={4}>用心呵护</Title>
                <Paragraph>
                  我们像对待自己的宠物一样照顾您的爱宠，提供贴心的关怀和照料。
                </Paragraph>
              </Card>
            </Col>
            <Col xs={24} sm={12} md={6}>
              <Card className="feature-card" bordered={false} style={{ textAlign: 'center' }}>
                <HomeOutlined style={{ fontSize: 48, color: '#1890ff', marginBottom: 16 }} />
                <Title level={4}>舒适环境</Title>
                <Paragraph>
                  提供宽敞、干净、舒适的寄养环境，让宠物在熟悉的环境中得到放松。
                </Paragraph>
              </Card>
            </Col>
            <Col xs={24} sm={12} md={6}>
              <Card className="feature-card" bordered={false} style={{ textAlign: 'center' }}>
                <SafetyOutlined style={{ fontSize: 48, color: '#52c41a', marginBottom: 16 }} />
                <Title level={4}>安全保障</Title>
                <Paragraph>
                  24小时监控和专业人员值守，确保您的宠物在寄养期间的安全。
                </Paragraph>
              </Card>
            </Col>
            <Col xs={24} sm={12} md={6}>
              <Card className="feature-card" bordered={false} style={{ textAlign: 'center' }}>
                <CustomerServiceOutlined style={{ fontSize: 48, color: '#faad14', marginBottom: 16 }} />
                <Title level={4}>专业服务</Title>
                <Paragraph>
                  由经验丰富的宠物护理师提供专业的照顾，包括喂食、遛狗、洗澡等服务。
                </Paragraph>
              </Card>
            </Col>
          </Row>
        </div>

        {/* 统计数据 */}
        <div style={{ background: '#f5f5f5', padding: '50px 0' }}>
          <Row gutter={32} justify="center">
            <Col xs={24} sm={8}>
              <Card bordered={false}>
                <Statistic 
                  title="服务宠物" 
                  value={5000} 
                  prefix={<UserOutlined />} 
                  suffix="只" 
                />
              </Card>
            </Col>
            <Col xs={24} sm={8}>
              <Card bordered={false}>
                <Statistic 
                  title="合作伙伴" 
                  value={50} 
                  prefix={<ShopOutlined />} 
                  suffix="家" 
                />
              </Card>
            </Col>
            <Col xs={24} sm={8}>
              <Card bordered={false}>
                <Statistic 
                  title="满意评价" 
                  value={98.5} 
                  precision={1}
                  prefix={<CommentOutlined />} 
                  suffix="%" 
                />
              </Card>
            </Col>
          </Row>
        </div>

        {/* 推荐服务 */}
        <div style={{ background: '#fff', padding: '50px 0' }}>
          <Title level={2} style={{ textAlign: 'center', marginBottom: 30 }}>热门寄养服务</Title>
          <Row gutter={[24, 24]}>
            {featuredServices.map(service => (
              <Col xs={24} sm={12} md={6} key={service.id}>
                <Card
                  hoverable
                  cover={<img alt={service.name} src={service.imageUrl} />}
                  actions={[
                    <div key="rating"><StarOutlined /> {service.rating}</div>,
                    <div key="price">¥{service.price}/天</div>,
                    <Link to="/service" key="book">预订</Link>
                  ]}
                >
                  <Card.Meta
                    title={service.name}
                    description={
                      <>
                        <Paragraph ellipsis={{ rows: 2 }}>{service.description}</Paragraph>
                        <Text type="secondary">提供商: {service.provider}</Text>
                      </>
                    }
                  />
                </Card>
              </Col>
            ))}
          </Row>
          <div style={{ textAlign: 'center', marginTop: 30 }}>
            <Button type="primary" size="large" onClick={() => navigate('/service')}>
              查看全部服务
            </Button>
          </div>
        </div>

        {/* 用户评价 */}
        <div style={{ background: '#f5f5f5', padding: '50px 0' }}>
          <Title level={2} style={{ textAlign: 'center', marginBottom: 30 }}>用户评价</Title>
          <Row gutter={[24, 24]} justify="center">
            {testimonials.map(testimonial => (
              <Col xs={24} sm={8} key={testimonial.id}>
                <Card bordered={false} style={{ height: '100%' }}>
                  <div style={{ marginBottom: 16 }}>
                    {[...Array(testimonial.rating)].map((_, i) => (
                      <StarOutlined key={i} style={{ color: '#faad14', marginRight: 4 }} />
                    ))}
                  </div>
                  <Paragraph style={{ fontSize: 16 }}>"{testimonial.content}"</Paragraph>
                  <div style={{ display: 'flex', alignItems: 'center' }}>
                    <Avatar src={testimonial.avatar} />
                    <Text strong style={{ marginLeft: 8 }}>{testimonial.name}</Text>
                  </div>
                </Card>
              </Col>
            ))}
          </Row>
        </div>

        {/* 注册提示 */}
        <div style={{ 
          background: 'linear-gradient(to right, #1890ff, #52c41a)', 
          padding: '60px 0',
          textAlign: 'center',
          color: '#fff'
        }}>
          <Title level={2} style={{ color: '#fff', marginBottom: 16 }}>立即加入我们</Title>
          <Paragraph style={{ fontSize: 18, marginBottom: 24 }}>
            注册成为会员，享受更多优质服务和专属优惠
          </Paragraph>
          {!currentUser && (
            <Button type="primary" size="large" ghost onClick={() => navigate('/register')}>
              立即注册
            </Button>
          )}
        </div>
      </Content>
      <Footer style={{ textAlign: 'center', background: '#001529', color: '#fff', padding: '24px 50px' }}>
        <div style={{ marginBottom: 16 }}>
          <Link to="/" style={{ color: '#fff', marginRight: 16 }}>首页</Link>
          <Link to="/service" style={{ color: '#fff', marginRight: 16 }}>服务</Link>
          <Link to="/about" style={{ color: '#fff', marginRight: 16 }}>关于我们</Link>
          <Link to="/contact" style={{ color: '#fff' }}>联系我们</Link>
        </div>
        <Divider style={{ borderColor: 'rgba(255,255,255,0.2)' }} />
        <div>
          宠物寄养系统 ©{new Date().getFullYear()} 版权所有
        </div>
      </Footer>
    </Layout>
  );
};

export default Home;