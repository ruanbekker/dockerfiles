import bcrypt
import time
import os
import logging
from random import choice, randint
from string import ascii_letters, punctuation, digits

logging.basicConfig(format='%(asctime)s %(levelname)s %(message)s', level=logging.INFO)

delay_time = os.environ.get('DELAY_TIME', '0')

def generate_hash(log_rounds):
    string_format = ascii_letters + punctuation + digits
    generated_string = "".join(choice(string_format) for x in range(randint(5, 20)))
    bcrypt_hash = bcrypt.hashpw(generated_string, bcrypt.gensalt(log_rounds))
    return bcrypt_hash

count = 0
log_rounds = choice([12,14,15])
while True:
    try:
        if count != 20:
            count = count+1
            gen_hash = generate_hash(log_rounds)
            logging.info("Your hash with {0} log rounds: {1}".format(log_rounds, gen_hash))
        else:
            logging.info("resetting")
            time.sleep(int(delay_time))
            count = 0
            log_rounds = choice([12,14,15])
    except:
        logging.error("something went wrong")
        time.sleep(10)
