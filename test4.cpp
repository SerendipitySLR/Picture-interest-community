  /*
 *File name:区间树查找 
 *Description:创建区间树，进行查找操作 
 */ 	
#include <iostream>
#include <string>
#include <windows.h>
#include <fstream>
using namespace std;
#define SIZE 15 
struct Node
{
	int low;//低端点 
	int high;//高端点 
	int max;
	string color;//颜色 
	Node *pParent;//父结点 
	Node *pLeft;//左孩子 
	Node *pRight;//右孩子 
};
 
class RBT
{
public:
	RBT();
	~RBT();
	void LeftRotate(Node* px);//左旋
	void RightRotate(Node* px);//右旋
	void Insert(Node* pz);//插入
	void InsertFixUp(Node* pz);//插入调整
	void InorderTreeWalk( Node* px );//中序遍历
	void IntervalSearch1( Node* x,Node* i );//递归的区间树.查找 
	bool IsOverlap(Node* x,Node* i);	//判断是否重叠 
	Node* GetRoot(){ return this->pT_root; }
	Node* GetNil(){ return this->pT_nil; }
	Node* IntervalSearch( Node* i );//区间树查找 
private:
	Node* pT_nil;
	Node* pT_root;
};
 
RBT::RBT()
{//构造一棵区间树 
	pT_nil = new Node; 
	pT_nil->color = "Black";//颜色设为BLACK 
	pT_nil->max = 0;
	pT_root = pT_nil;
}
RBT::~RBT()
{
	if( pT_nil != NULL )
		delete pT_nil;
}
 
void RBT::LeftRotate(Node *px)
{//左旋 
	Node* py = px->pRight;//用py记录px的右孩子 
	px->pRight = py->pLeft;//px的右孩子是py的左孩子 
	if( py->pLeft != pT_nil )
		py->pLeft->pParent = px;
	py->pParent = px->pParent;//py的父结点为px的父结点 
	if(px->pParent == pT_nil )//下面判断py为父结点的哪个孩子 
		pT_root = py;//根 
	else if( px == px->pParent->pLeft )
		px->pParent->pLeft = py;//左 
	else px->pParent->pRight = py;//右 
	py->pLeft = px;
	px->pParent = py;
	py->max = px->max;
	px->max = max( px->max,max(px->pLeft->max,px->pRight->max) );
}
void RBT::RightRotate(Node *py)
{//右旋 
	Node* px = py->pLeft;
	py->pLeft = px->pRight;
	px->pRight->pParent = py;
	px->pParent = py->pParent;
	if(py->pParent == pT_nil )
		pT_root = px;
	else if( py == py->pParent->pLeft )
		py->pParent->pLeft = px;
	else py->pParent->pRight = px;
	px->pRight = py;
	py->pParent = px;
	px->max = py->max;
	py->max = max( py->max,max(py->pLeft->max,py->pRight->max) );
}
 
void RBT::Insert( Node* pz)
{//插入 
	pz->max = pz->high;
	Node* py = pT_nil;
	Node *px  = pT_root;
	while( px != pT_nil )
	{
		px->max = max( pz->high,px->max );
		py = px;//用py记录px 
		if( pz->low < px->low )
			px = px->pLeft;
		else
			px = px->pRight;
	}
	pz->pParent = py;
	if( py == pT_nil )
		pT_root = pz;
	else if( pz->low < py->low )
		py->pLeft = pz;
	else
		py->pRight = pz;
	pz->pLeft = pT_nil;
	pz->pRight = pT_nil;
	pz->color = "Red";
	InsertFixUp( pz );
}
 
void RBT::InsertFixUp(Node* pz)
{//插入修正 
	Node* py = NULL;
	while( "Red" == pz->pParent->color )
	{
		if(pz->pParent == pz->pParent->pParent->pLeft )
		{
			py = pz->pParent->pParent->pRight;
			if( py->color == "Red" )
			{
				pz->pParent->color = "Black";
				py->color = "Black";
				pz->pParent->pParent->color = "Red";
				pz = pz->pParent->pParent;
			}
			else
			{
				if( pz == pz->pParent->pRight )
				{
					pz = pz->pParent;
					LeftRotate( pz );
				}
				pz->pParent->color = "Black";
				pz->pParent->pParent->color = "Red";
				RightRotate( pz->pParent->pParent );
			}
		}
		else if(pz->pParent == pz->pParent->pParent->pRight )
		{
			py = pz->pParent->pParent->pLeft;
			if( py->color == "Red" )
			{
				pz->pParent->color = "Black";
				py->color = "Black";
				pz->pParent->pParent->color = "Red";
				pz = pz->pParent->pParent;
			}
			else
			{
				if( pz == pz->pParent->pLeft )
				{
					pz = pz->pParent;
					RightRotate( pz );
				}
				pz->pParent->color = "Black";
				pz->pParent->pParent->color = "Red";
				LeftRotate( pz->pParent->pParent );
			}
		}
	}
	pT_root->color = "Black";
}
void RBT::InorderTreeWalk( Node* px )
{//中序遍历 
	if( px != pT_nil )
	{
		InorderTreeWalk( px->pLeft );
		cout << px->low << "-" << px->color << '-' << px->high <<endl;
		InorderTreeWalk( px->pRight );
	}
}
bool RBT::IsOverlap(Node* x,Node* i){
	//判断最小值
	if(x->low < i->low) {
		//x的开始范围小
		if(x->high >= i->low){
			return true;
		} else{
			return false;
		}
	}else if(x->low == i->low){
		return true;
	}else{
		//i的开始范围是最小
		if(i->high >= x->low){
			return true;
		}else{
			return false;
		} 
	}
}

//Node* RBT::IntervalSearch( Node* i)
//{//区间树查找 
//	Node* x = pT_root;//查找与i重叠的区间x的过程从以x为根的树根开始  
//	while( x != pT_nil && ( x->high < i->low || i->high < x->low ) )
//	{//当x指向pT.nil或找到重叠区间时过程结束 
//		if( x->pLeft != pT_nil && x->pLeft->max >= i->low )
//			x = x->pLeft;//去左区间查找 
//		else
//			x = x->pRight;//去右区间查找 
//	}
//	return x;
//}

void RBT::IntervalSearch1(Node* x,Node* i){
	//递归区间树查找 
//	Node* x = pT_root;//查找与i重叠的区间x的过程从以x为根的树根开始 
	if(IsOverlap(x,i))
	 cout << x->low << "-" << x->color << '-' << x->high <<',' << endl;
	if(x->pLeft != pT_nil && (x->pLeft->max >= i->low)) {
		//cout<<"i am left!"<<endl;
		IntervalSearch1(x->pLeft,i);
	}
	
	if((x->pRight != pT_nil&&(i->high>x->low))){
//		cout<<x->low<<','<<x->high<<"i am right!"<<endl;
	IntervalSearch1(x->pRight,i);
	}
}
//x->pLeft->max >= i->low && x->pRight->low <= i->high
int main()
{
	RBT rbt;
	fstream file("C:\\Users\\27587\\Desktop\\insert.txt");
	int A[8][2];		//接收insert文件中的数据
	int datalen=0;		//记录数据的长度 
	file>>datalen;
	//文件读取进入 
	for(int i=0;!file.eof();i++) {
		for(int j =0;j<2;j++){
			file>>A[i][j];
		}
	}
	//cout<<datalen<<endl;
//	for(int i=0;i<datalen;i++) {
//		for(int j =0;j<2;j++){
//			cout<<A[i][j]<<"\t";
//		}
//		cout<<endl;
//	}
	Node* ptemp = new Node[ SIZE ];
	for(int i=0;i<datalen;i++)
	{
		ptemp[i].low = A[i][0];
		ptemp[i].high = A[i][1];
		rbt.Insert( &ptemp[i] );
	}
 
	rbt.InorderTreeWalk( rbt.GetRoot() );
	cout << endl;
 
	bool bquit = true;//控制次数 
	Node temp;
	while(bquit)
	{
		cout << "输入区间: ";
		cin >> temp.low >> temp.high;
		rbt.IntervalSearch1(rbt.GetRoot(),&temp);
//		Node* p = rbt.IntervalSearch(&temp);
//		if(p != rbt.GetNil() )
//			cout << p->low << "-" << p->color << '-' << p->max <<',' << endl;
//		else
//			cout << "无重叠区间" << endl;
		cout << "1-继续/0-结束): ";
		cin >> bquit;
	}
	delete []ptemp;
	return 0;
}
