var stu = require('./controller.js');

module.exports = function(app){

  app.get('/get_all_stu', function(req, res){
    stu.get_all_stu(req, res);
  });

  app.get('/get_stu/:id', function(req, res){
    stu.get_stu(req, res);
  });

  app.get('/add_stu/:stu', function(req, res){
    stu.add_stu(req, res);
  });

}
