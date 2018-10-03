package config

var CfgElastic elasticConfig

func Init() {
	// // local
	// CfgElastic.Host = "http://localhost:9200"

	//dev / staging
	CfgElastic.Host = "http://devel-solr.tkpd:9200"
}
