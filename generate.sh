#!/bin/bash

DEFAULT_ROWS=1000000
NUM_ROWS=${1:-$DEFAULT_ROWS}
FILENAME="dummy_${NUM_ROWS}_rows.csv"

echo "Generating $NUM_ROWS rows..."
echo "ID,Name,Surname,Gender,Age,City,Score,Code,Status,Date,Email,Country,Verified,Phone,IP,Browser,OS,Device,Source,Plan,Referral" > "$FILENAME"

awk -v rows="$NUM_ROWS" '
BEGIN {
  srand();

  split("Alice,Bob,Charlie,David,Eve,Frank,Grace,Hannah,Ian,Julia,Sam,Manuel,Francisco,Fred,Carlos", names, ",");
  split("Smith,Johnson,Williams,Brown,Jones,Garcia,Miller,Davis,Martinez,Lopez,Marshall,Andres", surnames, ",");
  split("Male,Female,Other", genders, ",");
  split("New York,London,Berlin,Barcelona,Tokyo,Sydney,Paris,Dubai,Rome,Madrid,Beijing", cities, ",");
  split("USA,UK,Germany,Japan,Australia,France,UAE,Italy,Spain,China", countries, ",");
  split("example.com,test.org,demo.net,sample.io", domains, ",");
  split("Chrome,Firefox,Safari,Edge,Opera", browsers, ",");
  split("Windows,macOS,Linux,Android,iOS", oss, ",");
  split("Desktop,Laptop,Tablet,Mobile", devices, ",");
  split("Google,Facebook,Twitter,LinkedIn,Direct,Email", sources, ",");
  split("Free,Basic,Pro,Enterprise", plans, ",");

  for(i=1;i<=rows;i++) {
    name = names[int(rand()*length(names))+1];
    surname = surnames[int(rand()*length(surnames))+1];
    gender = genders[int(rand()*length(genders))+1];
    age = 18 + int(rand()*63);
    city = cities[int(rand()*length(cities))+1];
    score = sprintf("%.2f", rand()*100);
    code = int(rand()*99999) + 10000;
    status = (rand() > 0.7) ? "Suspended" : ((rand() > 0.5) ? "Active" : "Inactive");
    day = sprintf("%02d", 1 + int(rand()*28));
    date = "2025-" sprintf("%02d", 1 + int(rand()*12)) "-" day;
    email = tolower(name) i "@" domains[int(rand()*length(domains))+1];
    country = countries[int(rand()*length(countries))+1];
    verified = (rand() > 0.5) ? "true" : "false";
    phone = sprintf("+1-%03d-%03d-%04d", int(rand()*900+100), int(rand()*900+100), int(rand()*10000));
    ip = int(rand()*255) "." int(rand()*255) "." int(rand()*255) "." int(rand()*255);
    browser = browsers[int(rand()*length(browsers))+1];
    os = oss[int(rand()*length(oss))+1];
    device = devices[int(rand()*length(devices))+1];
    source = sources[int(rand()*length(sources))+1];
    plan = plans[int(rand()*length(plans))+1];
    referral = sprintf("REF_%04d", int(rand()*10000));

    printf "%d,%s,%s,%s,%d,%s,%s,%d,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s\n",
           i, name, surname, gender, age, city, score, code, status, date, email, country, verified,
           phone, ip, browser, os, device, source, plan, referral;
  }
}' >> "$FILENAME"

echo "CSV file with $NUM_ROWS rows generated: $FILENAME"

FILE_SIZE=$(du -h "$FILENAME" | cut -f1)
echo "File size: $FILE_SIZE"