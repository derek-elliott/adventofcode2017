import re
import json


def get_data():
    raw_data = []
    with open('data/dayseven.txt', 'r') as f:
        for line in f:
            raw_data.append(line.strip('\n'))
    return raw_data


def clean_data(raw_data):
    pattern = re.compile(r'([a-zA-Z]+)\s\(([0-9]+)\)\s?\-*>*\s?([a-z, ]*)')
    cleaned_data = []
    for line in raw_data:
        result = re.findall(pattern, line)
        if result[0][2] == '':
            cleaned_data.append(
                {'name': result[0][0], 'weight': int(result[0][1])})
        else:
            cleaned_data.append({'name': result[0][0], 'weight': int(result[0][1]), 'children': [
                                i for i in result[0][2].replace(',', '').split(' ')]})
    return cleaned_data


def write_data(cleaned_data):
    with open('data/cleaned_dayseven.json', 'w+') as f:
        json.dump(cleaned_data, f)


if __name__ == '__main__':
    raw_data = get_data()
    cleaned_data = clean_data(raw_data)
    write_data(cleaned_data)
