# End points

//if send just feature=2 or category=2 it returns products which feature id equals 2                    </br>
//feature=2&&category=2 it returns products which feature id equals 2 and category id equals 2          </br>
//selled=2&&category=2 it return products which selled id equals 2 and category id equals 2             </br>
//deleted=2&&feature=2 it return products which deleted id equals 2 and feature id equals 2             </br>
//all parameters = selled=productId && deleted=productId && feature=featureId && category=categoryId    </br>
r.GET("/stocks/filter", product.Filter)                                                                 </br>
