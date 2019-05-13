FROM docker.elastic.co/elasticsearch/elasticsearch:6.6.2
RUN bin/elasticsearch-plugin install analysis-icu

COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY config/elasticsearch.yml config/elasticsearch.yml
COPY config/setup.sh config/setup.sh

COPY templates/elasticsearch-mapping.json config/me.json
COPY templates/installed-software-tree-mapping.json config/is.json
COPY templates/os-tree-mapping.json config/os.json

RUN mkdir utils
COPY config/wait-for-it.sh utils/wait-for-it.sh

USER root
RUN chmod +x /docker-entrypoint.sh utils/wait-for-it.sh config/setup.sh
RUN chown -R elasticsearch:elasticsearch /docker-entrypoint.sh utils/wait-for-it.sh config/setup.sh

USER elasticsearch
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/usr/share/elasticsearch/bin/elasticsearch"]