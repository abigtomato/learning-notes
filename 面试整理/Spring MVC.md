# Spring MVC

## 基本概念

**概念**：Spring MVC是一个基于Java的实现了MVC设计模式的请求驱动类型的轻量级Web框架。其通过把模型-视图-控制器分离，将Web层进行职责解耦，把复杂的Web应用处理分为逻辑独立的几个部分。简化开发流程、减少出错、方便开发人员之间的配合。

**MVC**：是一种设计模式，即模型（Model）-视图（View）-控制器（Controller）三层架构的设计模式。用于实现前端页面展示与后端业务数据处理的分离。

* 分层设计，实现了业务系统各个组件之间的解耦，有利于业务系统的可扩展性和可维护性；
* 有利于系统的并行开发，提升开发效率。

**优点**：

* 支持各种视图技术，而不仅仅局限于JSP；
* 与Spring框架的天然集成；
* 清晰的角色分配：前对控制器、请求处理器映射、处理器适配器和视图解析器；
* 支持各种请求资源的映射策略。



## 核心组件

* **前端控制器（DispatcherServlet）**：作用是接收请求和响应结果。相当于各组件间的转发器，减少组件间的耦合度；
* **处理器映射器（HandlerMapping）**：作用是根据请求的URL查找对应的Handler；
* **处理器适配器（HandlerAdapter）**：适配器定义编写规则并执行开发人员编写的具体Handler；
* **处理器（Handler）**：开发人员编写的具体处理逻辑；
* **视图解析器（ViewResolver）**：进行视图的解析，根据逻辑名查找实际的视图；
* **视图（View）**：渲染视图，其实现类可以由不同的类型，如JSP、Freemarker等。



## 工作流程

* 客户端/浏览器发送请求到前端控制器DispatcherServlet；

* 前端控制器调用处理器映射器HandlerMapping，请求获取Handler；
* 处理器映射器根据请求的URL获取具体的处理器，生成处理器对象以及处理器拦截器，并返回给前端控制器；
* 前端控制器调用处理器适配器HandlerAdapter，请求执行Handler；
* 处理器适配器通过适配调用具体的处理器Handler执行业务逻辑；
* 处理器执行完毕后返回ModelAndView对象，该对象包含模型数据和视图名称；
* 处理器适配器将ModelAndView对象返回给前端控制器；
* 前端控制器将结果对象传递给视图解析器ViewResolver进行解析，即通过视图名称查找视图；
* 视图解析器解析完后返回具体的视图View对象；
* 前端控制器进行视图的渲染，即将模型数据填充进视图中；
* 返回渲染后的视图；
* 最后前端控制器将View响应给用户。

![img](assets/20180708224853769)



## 常用注解

### @Controller

在Spring MVC中，控制器Controller负责处理由前端控制器分发的请求，它将用户请求的数据经过业务处理层处理后封装成一个Model，然后再把该Model返回给对应的View进行展示。

被@Controller标记的类就是一个Spring MVC控制器对象。前端处理器将会扫描使用了该注解的类的方法，并检测方法是否被@RequestMapping修饰。

@Controller只是定义一个控制器类，而使用@RequestMapping注解的方法才是真正处理请求的Handler。

### @RequestMapping

该注解用于处理请求地址映射，可用于类或方法上。当用于类上时，表示类中的所有响应请求的方法都是以该地址做为前缀的。

### @ResponseBody

该注解用于将Controller的方法返回的对象，通过适当的HttpMessageConverter转换为指定格式后，写入到Response对象的body数据区中。当返回的数据不是HTML页面而是其他格式的数据时使用（如：JSON、XML格式）。

### @PathVariable和@RequestParam的区别

当参数被拼接到请求路径中传递时可以通过@PathVariable获取。而@RequestParam是用于获取HTTP请求中携带的参数。