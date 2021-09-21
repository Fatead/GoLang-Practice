package dao

import (
	"awesomeProject/src/bookstore/model"
	"awesomeProject/src/bookstore/utils"
	"fmt"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(count, amount, title, author, price, img_path, order_id) values(?,?,?,?,?,?,?)"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常，err:", err)
		return err
	}

	_, errExec := stmt.Exec(orderItem.Count, orderItem.Amount, orderItem.Title,
		orderItem.Author, orderItem.Price, orderItem.ImaPath, orderItem.OrderId)
	if errExec != nil {
		fmt.Println("添加订单项 出现异常，err:", errExec)
		return errExec
	}
	return nil
}

func FindOrderInfo(orderId string) ([]*model.OrderItem, error) {
	sqlStr := "select id, COUNT, amount, title, author, price, img_path, order_id from order_items where order_id = ? "
	rows, err := utils.Db.Query(sqlStr, orderId)
	if err != nil {
		fmt.Println("查询用户订单出现异常，err：", err)
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		err := rows.Scan(&orderItem.ID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImaPath, &orderId)
		if err != nil {
			fmt.Println("扫描写入出现异常，err:", err)
			return orderItems, err
		}
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
