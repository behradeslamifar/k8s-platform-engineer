# Kubernetes Platform Engineer
Build career to become Kubernetes Platform/Control Plane Engineer

For linkedin: Recently I have been rebuilding my deep Kubernetes platform engineering skills

## Phase 1: Programming (8 weeks)
Python Resources
- Python Crash Course - Third Edition
- Coding Interview Pattern

Golang Resources
- Go in action
- [Go Reference book in Farsi - Mostafa Solati](https://www.digikala.com/product/dkp-16392533/%DA%A9%D8%AA%D8%A7%D8%A8-%D9%85%D8%B1%D8%AC%D8%B9-%D8%A8%D8%B1%D9%86%D8%A7%D9%85%D9%87-%D9%86%D9%88%DB%8C%D8%B3%DB%8C-%D8%A8%D9%87-%D8%B2%D8%A8%D8%A7%D9%86-go-%D8%A7%D8%AB%D8%B1-%D9%85%D8%AD%D9%85%D8%AF%D9%85%D8%B5%D8%B7%D9%81%DB%8C-%D8%B5%D9%88%D9%84%D8%AA%DB%8C-%D9%87%D8%B4%D8%AC%DB%8C%D9%86-%D8%A7%D9%86%D8%AA%D8%B4%D8%A7%D8%B1%D8%A7%D8%AA-%D9%86%D8%B3%D9%84-%D8%B1%D9%88%D8%B4%D9%86)

<details><summary>Week 1-2</summary>
<p>

- Python Dev Environment in Linux
- Function, Method, Class and Module
- Simple Data Types (String, Number, Boolean)
- Data Strucures (Array/List, Graphs Hashtables, Linked Lists, Queues, Stacks, Trees)
- Algorithm (Prefix sum, ...)


</p>
</details>

## Phase 2: Operator Foundation (6-8 weeks)
Control Plane Engineering
- Go basics for controller pattern
- client-go
- Kubebuilder (build operator with Go)
- Build a real CRD
- implement reconciliation loop

<details><summary>Week 1-2</summary>
<p>
 
- Focuse on Go concurrency
  - [Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
- Channel
- [Context](https://go.dev/blog/context)
- Error Handling production-grade
 
</p>
</details>

<details><summary>Week 3-4</summary>
<p>
 
- [github.com: Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)
  - [Kubebuilder Book](https://book.kubebuilder.io/)
- Build Real CRD
- Reconciliation loop
- finalizer
- status management

Project: Operator 

</p>
</details>

<details><summary>Week 5-6</summary>
<p>
 
- Add metrics
- leader election
- failure simulation
- Upgrade CRD version
 
</p>
</details>

<details><summary>Week 7-8</summary>
<p>
 
</p>
</details>

## Phase 3:  Cluster Lifecycle Deep Dive (8~10 weeks)
Bare-Metal Mastery: Build Bare-metal cluster with `kubeadm` 
- kubeadm full lifecycle/bootstrap
- etcd failure/recovery backup/restore
- upgrade path/strategy
- node autoscaling simulation/node churn/node drain
- simulate failure/chaos drill.
- Cluster API

## Phase 4:  GPU / AI workload tie-in/exposure
Interesting Part
- LLM inference on K8s
- resource isolation
- scheduling experiments
- autoscaling GPU nodes
- Implement Reconciliation loop

