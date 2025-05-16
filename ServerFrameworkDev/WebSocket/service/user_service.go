package service

// import (
// 	"ServerFramework/global"
// 	"ServerFramework/model"
// 	"ServerFramework/model/request"
// 	"errors"

// 	"gorm.io/gorm"
// )

// // UpdateUser 根据请求信息更新用户信息
// func (u *UserService) UpdateUser(req request.UserUpdateRequest) (*model.User, error) {
//     // 假设这里根据用户ID更新用户信息，实际中需要根据业务逻辑获取用户ID
//     // 这里简单模拟一个用户ID，实际使用时需要从请求中获取正确的用户ID
//     userID := uint(1) 

//     var user model.User
//     if err := global.DB.First(&user, userID).Error; err != nil {
//         if errors.Is(err, gorm.ErrRecordNotFound) {
//             return nil, global.ErrUserNotFound
//         }
//         return nil, err
//     }

//     // 更新用户信息
//     if req.Username != "" {
//         user.Username = req.Username
//     }
//     if req.Email != "" {
//         user.Email = req.Email
//     }

//     if err := global.DB.Save(&user).Error; err != nil {
//         return nil, err
//     }

//     return &user, nil
// }
