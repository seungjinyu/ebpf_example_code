FROM apache/airflow:2.8.4

ADD requirements.txt .

RUN pip3 install apache-airflow==2.8.4 -r requirements.txt

# RUN export PYTHONPATH=/home/airflow/dags:/opt/airflow/dags

# RUN mkdir -p /opt/airflow/data/paper