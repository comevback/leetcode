# System Design Template

**contributed by [chaitraboggaram](https://leetcode.com/discuss/interview-question/system-design/5073436/System-Design-Template)**

### **1. Feature Expectations [5 min]**

```
	(1) Use Cases
	(2) Scenarios That Will Not Be Covered
	(3) Who Will Use
	(4) How Many Will Use
	(5) Usage Patterns
```

### **2. Estimations [5 min]**

```
	(1) Throughput
		  - Queries per second (QPS) for read and write queries.
	(2) Latency
		  - Expected latency for read and write queries.
	(3) Read/Write Ratio
	(4) Traffic Estimates
		  - Write (QPS, volume of data).
		  - Read (QPS, volume of data).
	(5) Storage Estimates
	(6) Memory Estimates
		  - If we are using a cache, what kind of data do we want to store in the cache?
		  - How much RAM and how many machines are needed?
		  - Amount of data to store on disk/SSD.
```

### **3. Design Goals [5 min]**

```
	(1) Latency and Throughput Requirements
	(2) Consistency vs. Availability
		  - Weak/strong/eventual consistency.
		  - Failover/replication for availability.
```

### **4. High-Level Design [5-8 min]**

```
	(1) APIs for Read/Write Scenarios for Crucial Components
	(2) Database Schema
	(3) Basic Algorithm
	(4) High-Level Design for Read/Write Scenario
```

### **5. Deep Dive [10-12 min]**

```
	(1) Scaling the Algorithm
	(2) Scaling Individual Components
		  - Availability, Consistency, and Scale story for each component
		  - Consistency and availability patterns.
	(3) Components to Consider
		  a) DNS
		  b) CDN (Push vs. Pull)
		  c) Load Balancers (Active-Passive, Active-Active, Layer 4, Layer 7)
		  d) Reverse Proxy
		  e) Application Layer Scaling (Microservices, Service Discovery)
		  f) Database options
				- RDBMS: ACID Properties, Primary-Secondary, Primary-Primary, Federation, Sharding, Denormalization, SQL Tuning - Postgres
					- Use-cases: Structured data with relationships
				- NoSQL: Key-Value, Wide-Column, Document - MongoDB, DynamoDB
					- Use-cases: Unstructured or semi-structured data
				- Graph: Neo4j, Amazon Neptune
					- Use-cases: Social networks, knowledge graphs, recommendation systems, and bioinformatics
				- NewSQL: Key-Value with ACID Properties - CockroachDB, Google Spanner, VoltDB
					- Use-cases: Transaction processing, real-time analytics and IoT device data
				- Time Series: Time-stamped data points - InfluxDB, TimescaleDB, Prometheus
					- Use-cases: IoT sensor data, financial market data, system metrics, and logs
				- Vector: High-dimensional vector data - Pinecone, Weaviate, KDB.AI
					- Use-cases: Machine learning, similarity search, and recommendation systems
				- Fast lookups:
					  - RAM (Bounded size) => Redis, Memcached.
					  - AP (Unbounded size) => Cassandra, RIAK, Voldemort, DynamoDB (default mode)
					  - CP (Unbounded size) => HBase, MongoDB, Couchbase, DynamoDB (consistent read setting).
		  g) Caches
				- Client caching, CDN caching, Webserver caching, Database caching, Application caching, Cache at query level, Cache at object level.
				- Eviction policies:
					  - Cache aside
					  - Write through
					  - Write behind
					  - Refresh ahead.
		  h) Asynchronism
				- Message queues
				- Task queues
				- Back pressure.
		  i) Communication
				- TCP
				- UDP
				- REST
				- RPC
				- WebSockets
```

### **6. Justify [5 min]**

```
	(1) Throughput of Each Layer
	(2) Latency Caused Between Each Layer
	(3) Overall Latency Justification
```

### **7. Key Metrics to Measure [3 min]**

```
	(1) Identify key metrics relevant to your system's design
			- Availability
			- Latency
			- Throttling
			- Request Patterns/Volume
			- Measure customer experience/feature specific metrics

	(2) Define metrics for infrastructure and resources
			- Tools like Graphana with Prometheus, AppDynamics, etc., can be used.
```

### **8. System Health Monitoring [2 min]**

```
	(1) Measure app index and latency of microservices
	(2) Tools like New Relic, AppDynamics can be used
			-  Use Grafana with Prometheus or AppDynamics for monitoring
	(3) Canaries - to simulate customer experience and pro-actively detect service degradation
```

### **9. Log Systems [2 min]**

```
	(1) Implement tools to gather and visualize metrics
	(2) Collect and analyze logs with ELK (Elastic, Logstash, Kibana) or Splunk.
```

### **10. Security [2 min]**

```
	(1) Firewall, encryptions at rest and in transit
	(2) TLS
	(3) Authentication, Authorization (AuthN/Z)
	(4) Limited Egress/Ingress
	(5) Principle of least privilege
```