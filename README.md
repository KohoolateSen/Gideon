<p align="center">
  <img src="https://github.com/user-attachments/assets/3b4969d5-2e43-42db-8b28-dd9b80c5a9b8" alt="Designer(2)" height="400em" width="400em"/>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-f0f0f0?logo=Go&logoColor=00e5ff&logoSize=auto" alt="Language - Golang">
  <img src="https://img.shields.io/badge/Gideon-white?logoSize=auto&logo=github&logoColor=black" alt=" Gideon">
  <img src="https://img.shields.io/badge/Telegram_bot-white?logoSize=auto&logo=Telegram&logoColor=24A1DE" alt=" Telegram bot">
</p>

# Gideon
Welcome to Gideon, your personal newsreader Telegram bot, inspired by the intelligent AI assistant from The Flash series. Just like its namesake, Gideon is designed to keep you informed with the latest news, ensuring you never miss out on important updates.

# Features 
Currently, Gideon excels at gathering the latest and most
relevant news for you! Soon, it will offer even more features.

# Getting Started

**1. Create API Tokens**
- Obtain a Telegram bot API token by creating a bot using BotFather.
- Generate a NewsAPI API token from https://newsapi.org.

**2. Clone & Configure**
- Clone the project
 ```
 git clone https://github.com/10xComrade/Gideon.git
 ```
- Change directory:
 ```
 cd Gideon
 ```
- Open the configuration file with a text editor:
 ```
 nano ./internal/config/config.yaml
 ```
- Insert your API tokens in the appropriate places.
- Optionally, adjust other configuration parameters like proxies.
- Save your changes.

**3. Start the Bot**
```
go run cmd/main.go
```

**4. Use the Bot Commands**
- Available commands:
  - `/news`
  - `/start`
  - `/help`  
    
- Command format:
```
/news <SUBJECT>
```

- Examples:
```
/news bitcoin
```
```
/news tesla
```

# InlineMode 
You can also access the latest news articles in any chat
without adding the bot—just use it inline.

- Example :
```
@RealGideonBot news@tesla
``` 
then choose & send your desired article.

# Limitations & Notes 

 - **Note 1 :**
   Once you use the /news command,
   you can seamlessly navigate between articles
   or read them using inline keyboard buttons.

 - **Note 2 :**
   Due to Telegram messaging rules,
   bots can send a maximum of 4091 characters per message.
   Longer articles can be read from their origin websites !

 - **Note 3 :**
   Gideon gathers news using a scraping method. Currently,
   we primarily extract content from `p` HTML tags,
   which may include gibberish and extra tags.
   Future updates will address this issue for improved accuracy.

  - **Note 4:**
  Due to current Telegram bot limitations, navigation and reading buttons are not available in inline mode.
  To read an article, you first need to select and send it via inline mode,
  then open the article's webpage directly within Telegram.

# Contributions 
If you encounter any problems or bugs with the project
or if you are interested in contributing,
please open an issue to let me know about your idea,
solution, or the problem you encountered.
