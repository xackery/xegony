{{template "header" .}}

{{template "navigation" .}}


<!-- Main Wrapper -->
<div id="wrapper">

<div class="content">

<div class="row">
    <div class="col-lg-4">
        <div class="hpanel forum-box">
            <div class="panel-heading">
                <span class="f"><a href="/character">Character</a> > {{.Character.Title}}</span>
            </div>
            <div class="panel-body">
                <div class="pull-right text-right">
                    <div class="btn-group">
                        <i class="fa fa-facebook btn btn-default btn-xs"></i>
                        <i class="fa fa-twitter btn btn-default btn-xs"></i>
                        <i class="fa fa-linkedin btn btn-default btn-xs"></i>
                    </div>
                </div>
                <img alt="logo" class="img-circle m-b m-t-md" src="/images/profile.jpg">
                <h3><a href="/profile.html">{{.Character.Title}} {{.Character.Name}} {{.Character.LastName}}</a></h3>
                <div class="text-muted font-bold m-b-xs">{{.Character.ZoneId}}</div>
                <p>
                   Level {{.Character.Level}} {{.Character.ClassName}}
                </p>
                <div class="progress m-t-xs full progress-small">
                    <div style="width: 65%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="65" role="progressbar" class=" progress-bar progress-bar-success">
                        <span class="sr-only">35% Complete (success)</span>
                    </div>
                </div>
            </div>
            <div class="border-right border-left">
                <section id="map">
                    <div id="map1" style="height: 200px"></div>
                </section>
            </div>
            <div class="panel-body">
                <dl>
                    <dt>Description lists</dt>
                    <dd>A description list is perfect for defining terms.</dd>
                    <dt>Euismod</dt>
                    <dd>Vestibulum id ligula porta felis euismod semper eget lacinia odio sem nec elit.</dd>
                    <dd>Donec id elit non mi porta gravida at eget metus.</dd>
                    <dt>Malesuada porta</dt>
                    <dd>Etiam porta sem malesuada magna mollis euismod.</dd>
                </dl>
            </div>
            <div class="panel-footer contact-footer">
                <div class="row">
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Projects: </span> <strong>200</strong></div>
                    </div>
                    <div class="col-md-4 border-right">
                        <div class="contact-stat"><span>Messages: </span> <strong>300</strong></div>
                    </div>
                    <div class="col-md-4">
                        <div class="contact-stat"><span>Views: </span> <strong>400</strong></div>
                    </div>
                </div>
            </div>

        </div>
    </div>
    <div class="col-lg-8">
        <div class="hpanel">
            <div class="hpanel">

            <ul class="nav nav-tabs">
                <li class="active"><a data-toggle="tab" href="/profile.html#tab-1">Project</a></li>
                <li class=""><a data-toggle="tab" href="/profile.html#tab-2">Messages</a></li>
            </ul>
            <div class="tab-content">
                <div id="tab-1" class="tab-pane active">
                    <div class="panel-body">
                        <strong>Lorem ipsum dolor sit amet, consectetuer adipiscing</strong>

                        <p>A wonderful serenity has taken possession of my entire soul, like these sweet mornings of spring which I enjoy with my whole heart. I am alone, and feel the charm of
                            existence in this spot, which was created for the bliss of souls like mine.</p>

                        <div class="table-responsive">
                            <table class="table table-striped">
                                <thead>
                                <tr>

                                    <th>#</th>
                                    <th>Project </th>
                                    <th>Name </th>
                                    <th>Phone </th>
                                    <th>Company </th>
                                    <th>Completed </th>
                                    <th>Task</th>
                                    <th>Date</th>
                                    <th>Action</th>
                                </tr>
                                </thead>
                                <tbody>
                                <tr>
                                    <td>1</td>
                                    <td>Project <small>This is example of project</small></td>
                                    <td>Patrick Smith</td>
                                    <td>0800 051213</td>
                                    <td>Inceptos Hymenaeos Ltd</td>
                                    <td><span class="pie">2/45</span></td>
                                    <td>20%</td>
                                    <td>Jul 14, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>2</td>
                                    <td>Alpha project</td>
                                    <td>Alice Jackson</td>
                                    <td>0500 780909</td>
                                    <td>Nec Euismod In Company</td>
                                    <td><span class="pie">1/5</span></td>
                                    <td>40%</td>
                                    <td>Jul 16, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>3</td>
                                    <td>Betha project</td>
                                    <td>John Smith</td>
                                    <td>0800 1111</td>
                                    <td>Erat Volutpat</td>
                                    <td><span class="pie">4/7</span></td>
                                    <td>75%</td>
                                    <td>Jul 18, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>4</td>
                                    <td>Gamma project</td>
                                    <td>Anna Jordan</td>
                                    <td>(016977) 0648</td>
                                    <td>Tellus Ltd</td>
                                    <td><span class="pie">12/3</span></td>
                                    <td>18%</td>
                                    <td>Jul 22, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>2</td>
                                    <td>Alpha project</td>
                                    <td>Alice Jackson</td>
                                    <td>0500 780909</td>
                                    <td>Nec Euismod In Company</td>
                                    <td><span class="pie">2/5</span></td>
                                    <td>40%</td>
                                    <td>Jul 16, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>1</td>
                                    <td>Project <small>This is example of project</small></td>
                                    <td>Patrick Smith</td>
                                    <td>0800 051213</td>
                                    <td>Inceptos Hymenaeos Ltd</td>
                                    <td><span class="pie">1/5</span></td>
                                    <td>20%</td>
                                    <td>Jul 14, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>4</td>
                                    <td>Gamma project</td>
                                    <td>Anna Jordan</td>
                                    <td>(016977) 0648</td>
                                    <td>Tellus Ltd</td>
                                    <td><span class="pie">2/8</span></td>
                                    <td>18%</td>
                                    <td>Jul 22, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>1</td>
                                    <td>Project <small>This is example of project</small></td>
                                    <td>Patrick Smith</td>
                                    <td>0800 051213</td>
                                    <td>Inceptos Hymenaeos Ltd</td>
                                    <td><span class="pie">15/5</span></td>
                                    <td>20%</td>
                                    <td>Jul 14, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>2</td>
                                    <td>Alpha project</td>
                                    <td>Alice Jackson</td>
                                    <td>0500 780909</td>
                                    <td>Nec Euismod In Company</td>
                                    <td><span class="pie">2/3</span></td>
                                    <td>40%</td>
                                    <td>Jul 16, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>3</td>
                                    <td>Betha project</td>
                                    <td>John Smith</td>
                                    <td>0800 1111</td>
                                    <td>Erat Volutpat</td>
                                    <td><span class="pie">4/5</span></td>
                                    <td>75%</td>
                                    <td>Jul 18, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>4</td>
                                    <td>Gamma project</td>
                                    <td>Anna Jordan</td>
                                    <td>(016977) 0648</td>
                                    <td>Tellus Ltd</td>
                                    <td><span class="pie">2/12</span></td>
                                    <td>18%</td>
                                    <td>Jul 22, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>2</td>
                                    <td>Alpha project</td>
                                    <td>Alice Jackson</td>
                                    <td>0500 780909</td>
                                    <td>Nec Euismod In Company</td>
                                    <td><span class="pie">2/3</span></td>
                                    <td>40%</td>
                                    <td>Jul 16, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>1</td>
                                    <td>Project <small>This is example of project</small></td>
                                    <td>Patrick Smith</td>
                                    <td>0800 051213</td>
                                    <td>Inceptos Hymenaeos Ltd</td>
                                    <td><span class="pie">1/5</span></td>
                                    <td>20%</td>
                                    <td>Jul 14, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                <tr>
                                    <td>4</td>
                                    <td>Gamma project</td>
                                    <td>Anna Jordan</td>
                                    <td>(016977) 0648</td>
                                    <td>Tellus Ltd</td>
                                    <td><span class="pie">10/50</span></td>
                                    <td>18%</td>
                                    <td>Jul 22, 2013</td>
                                    <td><a href="/profile.html#"><i class="fa fa-check text-success"></i></a></td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
                <div id="tab-2" class="tab-pane">
                    <div class="panel-body no-padding">


                        <div class="chat-discussion" style="height: auto">

                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a1.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Michael Smith </a>
                                    <span class="message-date"> Mon Jan 26 2015 - 18:39:23 </span>
                                            <span class="message-content">
											Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat.
                                            </span>

                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-success"><i class="fa fa-heart"></i> Love</a>
                                    </div>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a4.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Karl Jordan </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover.
                                            </span>
                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-default"><i class="fa fa-heart"></i> Love</a>
                                        <a class="btn btn-xs btn-primary"><i class="fa fa-pencil"></i> Message</a>
                                    </div>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a2.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Michael Smith </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration.
                                            </span>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a5.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Alice Jordan </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet.
                                                It uses a dictionary of over 200 Latin words.
                                            </span>
                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-warning"><i class="fa fa-eye"></i> Nudge</a>
                                    </div>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a6.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Mark Smith </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet.
                                                It uses a dictionary of over 200 Latin words.
                                            </span>
                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-success"><i class="fa fa-heart"></i> Love</a>
                                    </div>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a4.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Karl Jordan </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover.
                                            </span>
                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-default"><i class="fa fa-heart"></i> Love</a>
                                        <a class="btn btn-xs btn-primary"><i class="fa fa-pencil"></i> Message</a>
                                    </div>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a2.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Michael Smith </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration.
                                            </span>
                                </div>
                            </div>
                            <div class="chat-message">
                                <img class="message-avatar" src="/images/a5.jpg" alt="" >
                                <div class="message">
                                    <a class="message-author" href="/profile.html#"> Alice Jordan </a>
                                    <span class="message-date">  Fri Jan 25 2015 - 11:12:36 </span>
                                            <span class="message-content">
											All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet.
                                                It uses a dictionary of over 200 Latin words.
                                            </span>
                                    <div class="m-t-md">
                                        <a class="btn btn-xs btn-default"><i class="fa fa-thumbs-up"></i> Like </a>
                                        <a class="btn btn-xs btn-default"><i class="fa fa-heart"></i> Love</a>
                                    </div>
                                </div>
                            </div>

                        </div>

                    </div>
                </div>
            </div>


            </div>
        </div>
    </div>
</div>

</div>
    <!-- Right sidebar -->
    <div id="right-sidebar" class="animated fadeInRight">

        <div class="p-m">
            <button id="sidebar-close" class="right-sidebar-toggle sidebar-button btn btn-default m-b-md"><i class="pe pe-7s-close"></i>
            </button>
            <div>
                <span class="font-bold no-margins"> Analytics </span>
                <br>
                <small> Lorem Ipsum is simply dummy text of the printing simply all dummy text.</small>
            </div>
            <div class="row m-t-sm m-b-sm">
                <div class="col-lg-6">
                    <h3 class="no-margins font-extra-bold text-success">300,102</h3>

                    <div class="font-bold">98% <i class="fa fa-level-up text-success"></i></div>
                </div>
                <div class="col-lg-6">
                    <h3 class="no-margins font-extra-bold text-success">280,200</h3>

                    <div class="font-bold">98% <i class="fa fa-level-up text-success"></i></div>
                </div>
            </div>
            <div class="progress m-t-xs full progress-small">
                <div style="width: 25%" aria-valuemax="100" aria-valuemin="0" aria-valuenow="25" role="progressbar"
                     class=" progress-bar progress-bar-success">
                    <span class="sr-only">35% Complete (success)</span>
                </div>
            </div>
        </div>
        <div class="p-m bg-light border-bottom border-top">
            <span class="font-bold no-margins"> Social talks </span>
            <br>
            <small> Lorem Ipsum is simply dummy text of the printing simply all dummy text.</small>
            <div class="m-t-md">
                <div class="social-talk">
                    <div class="media social-profile clearfix">
                        <a class="pull-left">
                            <img src="/images/a1.jpg" alt="profile-picture">
                        </a>

                        <div class="media-body">
                            <span class="font-bold">John Novak</span>
                            <small class="text-muted">21.03.2015</small>
                            <div class="social-content small">
                                Injected humour, or randomised words which don't look even slightly believable.
                            </div>
                        </div>
                    </div>
                </div>
                <div class="social-talk">
                    <div class="media social-profile clearfix">
                        <a class="pull-left">
                            <img src="/images/a3.jpg" alt="profile-picture">
                        </a>

                        <div class="media-body">
                            <span class="font-bold">Mark Smith</span>
                            <small class="text-muted">14.04.2015</small>
                            <div class="social-content">
                                Many desktop publishing packages and web page editors.
                            </div>
                        </div>
                    </div>
                </div>
                <div class="social-talk">
                    <div class="media social-profile clearfix">
                        <a class="pull-left">
                            <img src="/images/a4.jpg" alt="profile-picture">
                        </a>

                        <div class="media-body">
                            <span class="font-bold">Marica Morgan</span>
                            <small class="text-muted">21.03.2015</small>

                            <div class="social-content">
                                There are many variations of passages of Lorem Ipsum available, but the majority have
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="p-m">
            <span class="font-bold no-margins"> Sales in last week </span>
            <div class="m-t-xs">
                <div class="row">
                    <div class="col-xs-6">
                        <small>Today</small>
                        <h4 class="m-t-xs">$170,20 <i class="fa fa-level-up text-success"></i></h4>
                    </div>
                    <div class="col-xs-6">
                        <small>Last week</small>
                        <h4 class="m-t-xs">$580,90 <i class="fa fa-level-up text-success"></i></h4>
                    </div>
                </div>
                <div class="row">
                    <div class="col-xs-6">
                        <small>Today</small>
                        <h4 class="m-t-xs">$620,20 <i class="fa fa-level-up text-success"></i></h4>
                    </div>
                    <div class="col-xs-6">
                        <small>Last week</small>
                        <h4 class="m-t-xs">$140,70 <i class="fa fa-level-up text-success"></i></h4>
                    </div>
                </div>
            </div>
            <small> Lorem Ipsum is simply dummy text of the printing simply all dummy text.
                Many desktop publishing packages and web page editors.
            </small>
        </div>

    </div>

    <!-- Footer-->
    <footer class="footer">
        <span class="pull-right">
            Example text
        </span>
        Company 2015-2020
    </footer>

</div>