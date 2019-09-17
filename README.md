# GoAPIs
Some APIs I made with Go as part of my internship at Optimus Information.

This repo has 2 APIs :-
	1. Express-API
	2. School-API
	
Express API is an ecommerce API in which you can add products to the website and users can select and add products to their carts. API has user authentication so each user can only view the products in his/her own cart.

School API, as its name suggests, is and API for a school management system. There are three categories of users in this API - students, teachers and the admins. The students can only access routes via the GET method. Teachers can access whatever the student can, and also add marks and notices. Meanwhile, the admin can view all the routes and add sections, subjects and update teacher and student details.
