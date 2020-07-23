// ----------------------------temp-------------------------------
type ReloanFacialTest struct{}

func (c *ReloanFacialTest) Cancel() {} // 暂时不需要

func (c *ReloanFacialTest) Start() {

	mqConf := mq.Config{
		Choose: beego.AppConfig.String("mq_choose"),
		QueueName: map[string]string{
			riskpush.QUEUE_RISK_RELOAN_IR: "text/json",
		},
		RabbitMQ: mq.RMQConfig{
			Server:          beego.AppConfig.String("mq_rabbitmq_risk_server"),
			Durable:         true,
			DeleteWhenUnuse: false,
			Exclusive:       false,
			NoWait:          false,
		},
	}

	if err := mq.Init(&mqConf); err != nil {
		panic(err)
	}

	var (
		wg   sync.WaitGroup
		sig  = make(chan os.Signal)
		stop = make(chan struct{}, 0)
	)
	wg.Add(1)

	go TaskReLoanFacialTest(&wg, stop) // 处理第二次用户人脸分数较低的任务

	go func() {
		signal.Notify(sig, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
		<-sig
		close(stop)
	}()

	wg.Wait()

}

func TaskReLoanFacialTest(wg *sync.WaitGroup, stop chan struct{}) {
	defer wg.Done()

	logs.Info("[TaskReLoanFacialTest] task start... ...")
	for {
		select {
		case <-stop:
			logs.Info("[TaskReLoanFacialTest] received stop signal, exit... ...")
			return
		default:

			var orderID int64 = 200718830200000686
			orderData, err := models.GetOrder(orderID)
			logs.Info("[TaskReLoanFacialTest] get order info: %#v", orderData)
			if err != nil || orderData.IsReloan != int(types.IsReloanYes) || orderData.CheckStatus != types.LoanStatus4Review {
				//orderData.RiskCtlStatus != types.RiskCtlReloanWaitFacialCompare2B
				logs.Error("[TaskReLoanFacialTest] 订单状态不正确,请检查. order: %#v, err: %v", orderData, err)
				time.Sleep(500 * time.Millisecond)
				continue
			}

			// 人脸对比流程
			f := &nubarium.FacialLog{
				OrderID:    orderData.Id,
				AccountID:  orderData.UserAccountId,
				IsReloan:   orderData.IsReloan == 1,
				FacialType: types.FacialCompareStep3C,
			}
			nubarium.FetchOriginInfo(f)

			return

		}
	}
}

// ----------------------------temp-------------------------------
