package language

func SetUp() {
	SetMessage("Gambling succeeded! %d has been paid.\nbalance: %d -> %d", "도박에 성공했습니다! %d원이 지급되었습니다.\n현재 잔액: %d원 -> %d원")

	SetMessage("Gambling failed! %d has been deducated.\nbalance: %d -> %d", "도박에 실패했습니다! %d원이 차감되었습니다.\n현재 잔액: %d원 -> %d원")

	SetMessage("Cannot enter money which is higher than your balance.\nbalance: %d", "현재 잔액보다 높은 돈은 입력할 수 없습니다.\n현재 잔액: %d원")

	SetMessage("%s's balance is %d.", "%s님의 현재 잔액은 %d원입니다.")

	SetMessage("Cannot pay baseMoney because you have money or coin.", "보유하신 돈이나 코인이 있어 기초자금을 지급할 수 없습니다.")

	SetMessage("%d was successfully paid!", "성공적으로 %d원이 지급되었습니다!")

	SetMessage("Already checked attendance today.", "오늘은 이미 출석체크를 했습니다.")

	SetMessage("Attendance check was successful! %d has been paid.\nbalance: %d -> %d", "출석체크에 성공했습니다! %d원이 지급되었습니다.\n현재 잔액: %d원 -> %d원")

	SetMessage("Already join in gambling", "이미 가입된 유저입니다.")

	SetMessage("Successfully join!", "성공적으로 가입되었습니다!")

	SetMessage("You are not join in gambling", "가입되지 않은 유저입니다.")
}
