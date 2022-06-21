import requests, logging
import google.cloud.logging

google.cloud.logging.Client().setup_logging()

# load tgbot credentials
tgbot_token = None
with open(".credentials", "r") as f:
    tgbot_token = f.read()
tg_chat_id = "-1001424398966"

def weekly_poll(request):
    try:
        url = "https://api.telegram.org/bot{}/sendPoll".format(tgbot_token)
        message = {
            "chat_id": tg_chat_id,
            "question": "SPX?",
            
        }
        response = requests.get(url, params=message)
        logging.info({
            "action": "ga weekly poll",
            "request_url": url,
            "request_body": message,
            "response_code": response.status_code,
            "response_text": response.text,
        })
        return ""
    except Exception as e:
        import traceback
        logging.error(''.join(traceback.format_exception(type(e), e, e.__traceback__)))
        return ""