// pages/order/confirm/index.js
import http from '../../../utils/http'

Page({

  // 页面的初始数据
  data: {
    goodsItem: [],
    totalPrice: 0.00,
    totalGoodsCount: 0,
    show: false,
    showView: true,
    addressList:[],
    message:"",
    address_option: [],
    address_value: {},
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getCartInfo()
    this.getDefaultAddress()
  },

  // 获取购物车信息
  async getCartInfo(){
    let res = await http.GET('/cart/query', {open_id: wx.getStorageSync('open_id')})
    this.setData({
      goodsItem: res.data.data.cart_item,
      totalPrice: res.data.data.total_price
    })
  },
  async getDefaultAddress(){
    let res = await http.GET('/address/list', {open_id: wx.getStorageSync('open_id')})
    let address_list = res.data.data;
    console.log("address_list", address_list)
    let address_option = []
    for (let index = 0; index < address_list.length; index++) {
      let element = address_list[index];
      let temp_str = element.detail + " " +element.name + " " + element.tel
      address_option.push({"text":temp_str, "value":element.id})
    }
    this.setData({
      address_option: address_option,
      address_value:address_list[1].id
    })
  },

  // 提交订单
  submitOrder(){
    this.setData({show: true})
  },

  // 备注改变
  MarkChange(res) {
    this.setData({
      message: res.detail
    })
  },

  // 确认支付
  confirmPay(){
    this.setData({showView: false})
    http.POST('/order/submit', {
      open_id: wx.getStorageSync('open_id'),
      message:this.data.message,
      address_id:this.data.address_value,
    })
  },

  // 返回首页
  backToHome(){
    wx.switchTab({url: '/pages/home/index'})
  }
})