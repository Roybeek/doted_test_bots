const {Telegraf} = require('telegraf')
const bot = new Telegraf('1958943268:AAFJkPyMNRN4cUVJhcVgwHIL3SnMMsH7Gn8')


let poll_flag = 0;
let name_surname = ''
bot.start(ctx => {
    poll_flag = 1
    name_surname = ''
    ctx.reply('Привет, введи свое имя')
    }
)

bot.on('message', ctx => {
    switch (poll_flag) {
        case 1:
            name_surname = ctx.message.text;
            poll_flag = 2;
            ctx.reply('Введи фамилию');
            break;
        case 2:
            poll_flag = 0;
            ctx.reply('Приятно познакомиться, ' + ctx.message.text + ' ' + name_surname);
            break;
    }
    }
)



bot.launch()