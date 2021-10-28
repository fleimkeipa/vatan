# End points

    //if send just feature=2 or category=2 it returns products which feature id equals 2                    
    //feature=2&&category=2 it returns products which feature id equals 2 and category id equals 2          
    //selled=2&&category=2 it return products which selled id equals 2 and category id equals 2             
    //deleted=2&&feature=2 it return products which deleted id equals 2 and feature id equals 2             
    //all parameters = selled=productId && deleted=productId && feature=featureId && category=categoryId    
    r.GET("/stocks/filter", product.Filter)                                                                 
