
---
date: "2025-01-29T00:00:00Z"
description: Master CAP Theorem for your software engineering interview with simple explanations, real-world examples, and PACELC insights. Learn trade-offs between consistency, availability, and partition tolerance.
header-img: "/assets/images/Distributed-System/CAP-Theorem.png"
keywords: CAP Theorem, PACELC Theorem, distributed systems, consistency, availability, partition tolerance, software engineering interview, system design, database trade-offs, interview preparation
tags:
- Distributed-Systems
- Interview-Questions
title: 'Interview Question: Explain CAP Theorem'
toc: true
---

## **1. What is the CAP Theorem?**  
The CAP Theorem states that a distributed system can only guarantee **two out of three properties** at the same time:  
1. **Consistency (C):** All users see the same data simultaneously.  
2. **Availability (A):** The system always responds to requests (even with stale data).  
3. **Partition Tolerance (P):** The system works even if parts of it lose communication.  

**Key Insight:** Network partitions (e.g., server crashes, dropped messages) are unavoidable. During a partition, you **must choose between C and A**.  

![Explain CAP Theorem](/assets/images/Distributed-System/CAP-Theorem.png)

## **2. Real-World Example: Rohit’s “Reminder” Startup**  
Imagine Rohit runs a company called *Reminder* where he writes down customer reminders in a diary. When he hires Raj to help, problems arise:  

### **Problem 1: Inconsistency**  
- **Scenario:** Rohit writes a reminder in his diary, but Raj doesn’t update his copy.  
- **Result:** Customers get inconsistent information.  
  ```  
  [Client] --> [Rohit's Diary: "Call X"]  
  [Client] --> [Raj's Diary: "No Entry"]  
  ```  
- **Solution:** Both must write entries simultaneously to ensure **consistency**.  

### **Problem 2: Availability Issues**  
- **Scenario:** If Raj is unavailable, Rohit can’t respond to requests until Raj returns.  
- **Result:** Delays frustrate customers.  
  ```  
  [Client Request] --> [Rohit: Available ✅]  
                     [Raj: Unavailable ❌]  
  ```  
- **Solution:** Accept entries even if one person is missing, and sync later.  

### **Problem 3: Network Partition**  
- **Scenario:** Rohit and Raj can’t communicate. The system must either:  
  - **Reject requests** (prioritize **consistency** but lose availability).  
  - **Accept requests** (prioritize **availability** but risk inconsistency).  
  ```  
  Partition Occurs:  
  [Rohit] --✂️ Network Failure ✂️-- [Raj]  
  Choose:  
  1. Reject writes ➔ Consistency ✔️, Availability ❌  
  2. Accept writes ➔ Availability ✔️, Consistency ❌  
  ```  

## **3. CAP Trade-offs: CP vs. AP vs. CA**  

| **System Type** | **Prioritizes** | **Sacrifices** | **Example** |  
|------------------|------------------|-----------------|-------------|  
| **CP** | Consistency + Partition Tolerance | Availability | Banking systems (e.g., PostgreSQL) |  
| **AP** | Availability + Partition Tolerance | Consistency | Social media (e.g., Cassandra) |  
| **CA** | Consistency + Availability | Partition Tolerance | Single-server databases (e.g., non-replicated MySQL) |  

⚠️ **CA systems are theoretical** – real-world distributed systems **must handle partitions (P)**.  

## **4. PACELC Theorem: Beyond CAP**  
The PACELC Theorem extends CAP to address **normal operations** (no partitions):  
- **PAC:** During a **partition (P)**, choose between **A** (availability) and **C** (consistency).  
- **ELC:** **Else (E)**, choose between **L** (latency) and **C** (consistency).  

### **Examples**  
- **Banking Apps:** Prioritize **consistency** (e.g., ATMs must show accurate balances).  
- **Social Media:** Prioritize **availability** (e.g., Facebook News Feed can show stale posts).  
- **Messaging Apps:** Prioritize **consistency** (e.g., WhatsApp must show messages in order).  

## **5. CAP in Popular Databases**  
| **Database** | **CAP Type** | **Use Case** |  
|--------------|--------------|--------------|  
| Cassandra | AP | High availability (social media, IoT) |  
| MongoDB | CP (tunable) | Consistency-focused apps (e-commerce) |  
| DynamoDB | AP (default) | Scalable web apps (Amazon) |  

## **Key Takeaways**  
✅ **CAP Theorem:** Choose 2/3 properties during a partition.  
✅ **PACELC:** Adds latency vs. consistency trade-offs in normal conditions.  
✅ **AP vs. CP:** Match the system to the use case (e.g., social media = AP, banking = CP).  

**Further Reading:**  
- [CAP Theorem Explained (BMC)](https://www.bmc.com/blogs/cap-theorem/)  
- [PACELC Deep Dive (Educative)](https://www.educative.io/blog/what-is-cap-theorem)