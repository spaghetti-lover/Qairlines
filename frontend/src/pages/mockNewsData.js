// mockNewsData.js - File quản lý mock data cho news

// Dữ liệu mẫu ban đầu
const initialMockNews = [
  {
    id: '1',
    title: '[VN102] Chuyến bay bị delay',
    description: 'Đấm "đọc thêm" để xem chi tiết',
    content: 'Do ảnh hưởng của thời tiết, chuyến bay mang số hiệu VN102 sẽ bị delay khoảng 2 tiếng. Ngành hàng không không chỉ đơn thuần là vận chuyển con người từ điểm A đến điểm B; nó còn là việc tạo ra một hành trình an toàn, hiệu quả và ngày càng thoải mái. Trong bối cảnh cạnh tranh gay gắt và kỳ vọng của hành khách ngày càng cao, các hãng hàng không liên tục đầu tư vào việc nâng cấp mọi khía cạnh của trải nghiệm bay. Từ những chiếc máy bay đầu tiên chỉ tập trung vào chức năng, đến nay, chúng ta đang chứng kiến sự bùng nổ của các tiện ích và công nghệ được thiết kế để tối ưu hóa sự hài lòng của hành khách.Hành khách vui lòng liên hệ quầy vé để biết thêm chi tiết.',
    image: '/QAirline-card.png',
    authorId: '1',
    authorName: 'John Wick',
    createdAt: '2024-01-15T10:30:00Z',
    updatedAt: null,
    status: 'published'
  },
  {
    id: '2',
    title: 'New Routes Announced',
    description: 'Giới thiệu tuyến bay mới',
    content: 'Chúng tôi rất vui mừng thông báo về việc mở rộng mạng lưới đường bay của mình với các tuyến bay mới từ Hà Nội đến Đà Nẵng và từ TP.HCM đến Nha Trang. Các chuyến bay sẽ bắt đầu hoạt động từ ngày 1 tháng 2 năm 2024.',
    image: '/tour-2.jpg',
    authorId: '1',
    authorName: 'Nguyễn Văn A',
    createdAt: '2024-01-14T15:45:00Z',
    updatedAt: null,
    status: 'published'
  },
  {
    id: '3',
    title: 'Special Discounts for Summer',
    description: 'Khuyến mãi đặc biệt mùa hè',
    content: 'Để chào mừng mùa hè, chúng tôi xin gửi đến quý khách hàng chương trình khuyến mãi đặc biệt với mức giảm giá lên đến 30% cho tất cả các chuyến bay nội địa. Chương trình áp dụng từ ngày 20 tháng 1 đến hết ngày 31 tháng 3 năm 2024.',
    image: '/tour-3.jpg',
    authorId: '1',
    authorName: 'Nguyễn Văn A',
    createdAt: '2024-01-13T08:20:00Z',
    updatedAt: null,
    status: 'published'
  },
];

// Mock user data
const mockUsers = {
  '1': {
    userId: '1',
    firstName: 'Admin',
    lastName: 'User',
    email: 'admin@example.com',
    role: 'admin'
  }
};

// Hàm random hình ảnh tour
const getRandomTourImage = () => {
  const tourNumber = Math.floor(Math.random() * 5) + 1; // Random số từ 1-5
  return `/tour-${tourNumber}.jpg`;
};

// Class quản lý mock data
class MockNewsDataService {
  constructor() {
    if (!MockNewsDataService.instance) {
      // Khởi tạo dữ liệu từ localStorage hoặc dữ liệu mặc định
      this.loadFromStorage();
      MockNewsDataService.instance = this;
    }
    return MockNewsDataService.instance;
  }

  // Tải dữ liệu từ localStorage (nếu có)
  loadFromStorage() {
    try {
      const storedNews = typeof window !== 'undefined' ? localStorage.getItem('mockNewsData') : null;
      const storedUsers = typeof window !== 'undefined' ? localStorage.getItem('mockUsersData') : null;
      
      this.news = storedNews ? JSON.parse(storedNews) : [...initialMockNews];
      this.users = storedUsers ? JSON.parse(storedUsers) : { ...mockUsers };
    } catch (error) {
      console.warn('Error loading from localStorage:', error);
      this.news = [...initialMockNews];
      this.users = { ...mockUsers };
    }
  }

  // Lưu dữ liệu vào localStorage
  saveToStorage() {
    try {
      if (typeof window !== 'undefined') {
        localStorage.setItem('mockNewsData', JSON.stringify(this.news));
        localStorage.setItem('mockUsersData', JSON.stringify(this.users));
      }
    } catch (error) {
      console.warn('Error saving to localStorage:', error);
    }
  }

  // Lấy tất cả bài viết
  async getAllNews() {
    return new Promise((resolve) => {
      setTimeout(() => {
        resolve({ 
          data: [...this.news].sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt)),
          total: this.news.length
        });
      }, 300); // Giả lập delay API
    });
  }

  // Lấy bài viết theo ID
  async getNewsById(id) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const newsItem = this.news.find(item => item.id === id);
        if (newsItem) {
          resolve({ data: { ...newsItem } });
        } else {
          reject(new Error('Không tìm thấy bài viết'));
        }
      }, 200);
    });
  }

  // Tạo bài viết mới
  async createNews(newsData) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        try {
          const newId = Date.now().toString();
          const currentUser = this.users[newsData.authorId] || this.users['1'];
          
          const newNews = {
            id: newId,
            title: newsData.title,
            description: newsData.description,
            content: newsData.content,
            image: getRandomTourImage(), // Luôn sử dụng hình ảnh random
            authorId: newsData.authorId,
            authorName: `${currentUser.firstName} ${currentUser.lastName}`,
            createdAt: new Date().toISOString(),
            updatedAt: null,
            status: 'published'
          };

          this.news.unshift(newNews);
          this.saveToStorage();
          
          resolve({ 
            data: newNews,
            message: 'Bài viết đã được tạo thành công'
          });
        } catch (error) {
          reject(new Error('Lỗi khi tạo bài viết: ' + error.message));
        }
      }, 500);
    });
  }

  // Cập nhật bài viết
  async updateNews(id, newsData) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        try {
          const index = this.news.findIndex(item => item.id === id);
          if (index !== -1) {
            const currentUser = this.users[newsData.authorId] || this.users['1'];
            
            this.news[index] = {
              ...this.news[index],
              title: newsData.title,
              description: newsData.description,
              content: newsData.content,
              image: getRandomTourImage(), // Luôn sử dụng hình ảnh random
              authorId: newsData.authorId,
              authorName: `${currentUser.firstName} ${currentUser.lastName}`,
              updatedAt: new Date().toISOString()
            };

            this.saveToStorage();
            
            resolve({ 
              data: this.news[index],
              message: 'Bài viết đã được cập nhật thành công'
            });
          } else {
            reject(new Error('Không tìm thấy bài viết để cập nhật'));
          }
        } catch (error) {
          reject(new Error('Lỗi khi cập nhật bài viết: ' + error.message));
        }
      }, 500);
    });
  }

  // Xóa bài viết
  async deleteNews(id) {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        try {
          const index = this.news.findIndex(item => item.id === id);
          if (index !== -1) {
            const deletedNews = this.news.splice(index, 1)[0];
            this.saveToStorage();
            
            resolve({ 
              data: deletedNews,
              message: 'Bài viết đã được xóa thành công'
            });
          } else {
            reject(new Error('Không tìm thấy bài viết để xóa'));
          }
        } catch (error) {
          reject(new Error('Lỗi khi xóa bài viết: ' + error.message));
        }
      }, 300);
    });
  }

  // Tìm kiếm bài viết
  async searchNews(query, filters = {}) {
    return new Promise((resolve) => {
      setTimeout(() => {
        let filteredNews = [...this.news];

        // Tìm kiếm theo từ khóa
        if (query) {
          const searchTerm = query.toLowerCase();
          filteredNews = filteredNews.filter(item =>
            item.title.toLowerCase().includes(searchTerm) ||
            item.description.toLowerCase().includes(searchTerm) ||
            item.content.toLowerCase().includes(searchTerm) ||
            item.authorName.toLowerCase().includes(searchTerm)
          );
        }

        // Lọc theo tác giả
        if (filters.authorId) {
          filteredNews = filteredNews.filter(item => item.authorId === filters.authorId);
        }

        // Lọc theo ngày
        if (filters.dateFrom) {
          filteredNews = filteredNews.filter(item => 
            new Date(item.createdAt) >= new Date(filters.dateFrom)
          );
        }

        if (filters.dateTo) {
          filteredNews = filteredNews.filter(item => 
            new Date(item.createdAt) <= new Date(filters.dateTo)
          );
        }

        // Sắp xếp
        filteredNews.sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt));

        resolve({ 
          data: filteredNews,
          total: filteredNews.length,
          query,
          filters
        });
      }, 300);
    });
  }

  // Lấy thông tin user hiện tại
  async getCurrentUser(userId = '1') {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const user = this.users[userId];
        if (user) {
          resolve({ data: user });
        } else {
          reject(new Error('Không tìm thấy thông tin người dùng'));
        }
      }, 200);
    });
  }

  // Lấy thống kê
  async getNewsStats() {
    return new Promise((resolve) => {
      setTimeout(() => {
        const today = new Date();
        const todayStr = today.toDateString();
        const thisWeek = new Date(today.getFullYear(), today.getMonth(), today.getDate() - 7);
        const thisMonth = new Date(today.getFullYear(), today.getMonth(), 1);

        const stats = {
          total: this.news.length,
          today: this.news.filter(item => 
            new Date(item.createdAt).toDateString() === todayStr
          ).length,
          thisWeek: this.news.filter(item => 
            new Date(item.createdAt) >= thisWeek
          ).length,
          thisMonth: this.news.filter(item => 
            new Date(item.createdAt) >= thisMonth
          ).length,
          byAuthor: this.news.reduce((acc, item) => {
            acc[item.authorName] = (acc[item.authorName] || 0) + 1;
            return acc;
          }, {}),
          recent: this.news
            .sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt))
            .slice(0, 5)
        };

        resolve({ data: stats });
      }, 200);
    });
  }

  // Reset dữ liệu về mặc định
  resetToDefault() {
    this.news = [...initialMockNews];
    this.users = { ...mockUsers };
    this.saveToStorage();
    return { message: 'Đã reset dữ liệu về mặc định' };
  }

  // Xuất dữ liệu
  exportData() {
    return {
      news: this.news,
      users: this.users,
      exportedAt: new Date().toISOString()
    };
  }

  // Nhập dữ liệu
  importData(data) {
    try {
      if (data.news && Array.isArray(data.news)) {
        this.news = data.news;
      }
      if (data.users && typeof data.users === 'object') {
        this.users = data.users;
      }
      this.saveToStorage();
      return { message: 'Nhập dữ liệu thành công' };
    } catch (error) {
      throw new Error('Lỗi khi nhập dữ liệu: ' + error.message);
    }
  }
}

// Export singleton instance
const mockNewsDataService = new MockNewsDataService();

export default mockNewsDataService;

// Export class để có thể tạo instance mới nếu cần
export { MockNewsDataService };