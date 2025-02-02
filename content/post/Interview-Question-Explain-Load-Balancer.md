
---
date: "2025-02-02T00:00:00Z"
description: Ace your software engineering interview with this comprehensive guide on load balancers. Learn core concepts, deployment strategies, algorithms, types, and key configurations.
header-img: "/assets/images/DB/Load-Balancer.png"
keywords: load balancer interview guide, scalability techniques, server health checks, reverse proxy, SSL offloading, round robin algorithm, AWS load balancer, DNS vs network load balancer, high availability systems, software engineering interview preparation
tags:
- Distributed-Systems
- Interview-Questions
title: 'Interview Question: Explain Load Balancer'
toc: true
---

## **What is a Load Balancer?**  
A **load balancer** acts like a traffic cop for your application, directing incoming requests across multiple servers to prevent overload and ensure smooth performance. It’s the entry point of your system, hiding server details and optimizing traffic flow.  
 
- **Placement**: Sits at the edge of data centers (like a gatekeeper).  
- **Functionality**:  
  - Distributes traffic to servers.  
  - Proxies requests to hide server IPs (e.g., clients only see the load balancer’s IP).  
  - Routes responses back to users.  

**Visual Analogy**: Imagine a supermarket with 10 cashiers. The load balancer ensures customers (requests) are evenly distributed to avoid long queues.  

![Explain Load Balancer](/assets/images/Distributed-System/Load-Balancer.png)

---

## **Why Use a Load Balancer?**  
### 1. **Scalability**  
- **Horizontal Scaling**: Add/remove servers like adding more cashiers during rush hour.  
- **Traffic Spikes**: Distribute sudden surges (e.g., Black Friday sales) without crashing servers.  

### 2. **Resilience**  
- **Health Checks**: Like a doctor checking pulse, the load balancer pings servers to ensure they’re alive.  
- **Auto-Adjustment**: Removes unhealthy servers and adds new ones during traffic spikes.  

### 3. **Higher Availability**  
- **Rolling Deployments**: Update servers one by one (e.g., deploy code to Server 1, test, then move to Server 2).  
- **Zero Downtime**: Users never notice updates.  

### 4. **Security**  
- **DDoS Protection**: Blocks malicious traffic before it reaches servers.  
- **IP Masking**: Hides server IPs, reducing attack surfaces.  

### 5. **Performance**  
- **SSL Offloading**: Handles encryption/decryption, freeing servers to focus on processing requests.  
- **Traffic Compression**: Sends smaller data packets for faster responses.  

---

## **Deployment Process with a Load Balancer**  
Here’s how companies deploy updates without downtime:  
1. **Detach a Server**: Temporarily remove Server 1 from the pool.  
2. **Deploy Code**: Update the application on Server 1.  
3. **Sanity Check**: Test if Server 1 works post-update.  
4. **Reattach**: Add Server 1 back to the pool.  
5. **Repeat**: Do this for all servers sequentially.  

**Visual Guide**:  
```
[Client] → [Load Balancer]  
           ↓         ↓  
        [Server 1 (Updated)]  [Server 2 (Old)] → ...  
```  

---

## **How Load Balancers Work**  
### **Reverse Proxy**  
- Accepts client requests → forwards to servers → returns responses.  
- **Example**: When you visit `example.com`, the load balancer (not the server) handles your request.  

### **Health Checks**  
- **Active Checks**: Regularly sends “Are you alive?” signals to servers.  
- **Recovery**: If Server 2 fails, traffic shifts to Server 1 and 3. Once Server 2 recovers, it rejoins the pool.  

### **Self-Scaling**  
- **AWS Example**: Automatically scales from 1 to 100 nodes during traffic surges.  

---

## **Load Balancing Algorithms**  
1. **Round Robin**:  
   - Sends requests sequentially: Server 1 → Server 2 → Server 3 → repeat.  
   - *Use Case*: Evenly distributed, stateless traffic (e.g., static websites).  

2. **Weighted Round Robin**:  
   - Assigns more requests to powerful servers (e.g., Server 1 handles 60% traffic).  

3. **Least Connections**:  
   - Directs traffic to the server with the fewest active requests.  

4. **IP Hash**:  
   - Uses client IP to assign a fixed server (useful for session persistence).  

```
Clients → [Load Balancer]  
           ↙   ↓   ↘  
        [S1] [S2] [S3]  
```  

---

## **Key Configuration Components**  
1. **Listeners**:  
   - Rules for traffic entry (e.g., HTTP on port 80 → Server port 8080).  

2. **Routing**:  
   - Conditional rules (e.g., send `/admin` requests to admin servers).  

3. **Target Groups**:  
   - Group servers by role (e.g., `static-content-servers`, `payment-servers`).  

---

## **Types of Load Balancers**  
1. **Application Load Balancer (ALB)**:  
   - **Layer 7 (HTTP/HTTPS)**: Routes based on URL paths (e.g., `/api` vs `/images`).  
   - *Example*: AWS ALB.  

2. **Network Load Balancer (NLB)**:  
   - **Layer 4 (TCP/UDP)**: Handles high-speed, low-latency traffic (e.g., gaming).  

3. **DNS Load Balancer**:  
   - Distributes traffic via DNS (e.g., returns different IPs for `example.com`).  

---

## **Interview Cheat Sheet**  
- **Explain Scalability**: “Load balancers let you add servers horizontally to handle traffic spikes.”  
- **Health Checks**: “They ensure only healthy servers get traffic, improving reliability.”  
- **Round Robin vs Least Connections**: “Round Robin cycles servers; Least Connections picks the least busy one.”  

**Common Questions**:  
1. **How does a load balancer improve security?**  
   - *Answer*: Hides server IPs, blocks DDoS attacks, and manages SSL.  

2. **What happens if a server fails during deployment?**  
   - *Answer*: The load balancer reroutes traffic to active servers.  
