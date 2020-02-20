package main

import (

    "fmt"
    "os"
	"os/signal"
    "github.com/bwmarrin/discordgo"
    "syscall"
    "math/rand"
    "sync"
    "time"
    "strconv"
)

const token string = "token"
var BotID string


func main() {
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
    sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
    return
}



var onlyOnce sync.Once

//prepare d6\20 dice
var diceTwenty = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
                      11, 12, 13, 14, 15, 16, 17, 18,
                      19, 20}
var diceSix = []int{1, 2, 3, 4, 5, 6}
var diceEight = []int{1, 2, 3, 4, 5, 6, 7, 8}

func rollDiceTwenty() int {
    
    onlyOnce.Do(func() {
        rand.Seed(time.Now().UnixNano()) //this runs once
    })
    
    return diceTwenty[rand.Intn(len(diceTwenty))]
}

func rollDiceSix() int {
    
    onlyOnce.Do(func() {
        rand.Seed(time.Now().UnixNano()) //this runs once
    })
    
    return diceSix[rand.Intn(len(diceSix))]
}

func rollDiceEight() int {
    
    onlyOnce.Do(func() {
        rand.Seed(time.Now().UnixNano()) //this runs once
    })
    
    return diceEight[rand.Intn(len(diceEight))]
}

func advrollDiceTwenty() int {
    
    onlyOnce.Do(func() {
        rand.Seed(time.Now().UnixNano()) //this runs once
    })
    
    return diceTwenty[rand.Intn(len(diceTwenty))]
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
 
    dice20 := rollDiceTwenty()
    d20 := strconv.Itoa(dice20)
    advdice20 := advrollDiceTwenty()
    advd20 := strconv.Itoa(advdice20)
    dice6 := rollDiceSix()
    d6 := strconv.Itoa(dice6)
    dice8 := rollDiceEight()
    d8 := strconv.Itoa(dice8)
    if m.Author.ID == BotID {
        return
    }
    
    if m.Content == "!roll" {
        _, _ = s.ChannelMessageSend("678667262437228544", "D20: " + d20)
    }
    
   if m.Content == "!rolld6" {
        _, _ = s.ChannelMessageSend("678667262437228544", "D6: " + d6)
    }
    
    if m.Content == "!advroll" {
        _, _ = s.ChannelMessageSend("678667262437228544", "D20: " + d20 + " D20: " + advd20)
    }
    
    if m.Content == "!rolld8" {
        _, _ = s.ChannelMessageSend("678667262437228544", "D8: " + d8)
    }
}
