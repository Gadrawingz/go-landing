 var dayOfWeek int = 3

   switch dayOfWeek {
      case 1:
         fmt.Println("It is Monday")
         fallthrough
      case 2:
         fmt.Println("It is Tuesday")
         fallthrough
      case 3:
         fmt.Println("It is Wednesday")
         fallthrough
      case 4:
         fmt.Println("It is Thursday")
         fallthrough
      case 5:
         fmt.Println("It is Friday")
         fallthrough
      case 6:
         fmt.Println("It is Saturday")
      case 7:
         fmt.Println("It is Sunday")
      default:
         fmt.Println("It\'s Invalid Day")
